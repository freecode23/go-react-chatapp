package chat

import (
	"fmt"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/freecode23/go-react-chatapp/pkg/message"
)

type Chatroom struct {
	chatroomName            string
	RegisteredClientsChan   chan *Client
	unregisteredClientsChan chan *Client
	clientsMap              map[*Client]bool // key client, value bool. made up set like in python
	messagesChan            chan message.Message
	cache                   cache.Cache
}

/*
*
Init new chatroomwhich is just a collection of clients that are connected together
Think of this as a single chat room
Return reference to a new Chatroom
*
*/
func NewChatroom(cacheObj cache.Cache, chatroomName string) *Chatroom {

	// return chatroom struct
	return &Chatroom{
		chatroomName:            chatroomName,
		RegisteredClientsChan:   make(chan *Client),
		unregisteredClientsChan: make(chan *Client),
		clientsMap:              make(map[*Client]bool),
		messagesChan:            make(chan message.Message),
		cache:                   cacheObj,
	}
}

/*
*
handled by go routine. there 1 go routine per Chatroom
it will handle any client request
register, unregister, broadcast
This is done continuously. as long as there is a client
coming in, go routine will do something
*
*/
func (cr *Chatroom) ProcessChatroomEvents() {

	for {
		select {

		// case1: read the client that register
		// client here looks like this: memory address &{ 0x140001542c0 0x1400006e0e0 <nil>}
		// does not print out the actual values stored in pointers
		case client := <-cr.RegisteredClientsChan:

			// 1. insert client
			cr.clientsMap[client] = true
			fmt.Println("chatroom: new client reg:", len(cr.clientsMap))

			for client := range cr.clientsMap {

				client.WsConn.WriteJSON(message.Message{

					Body: "New User Joined...",
				})
			}
			break

		// case2. read the client that unregister
		case client := <-cr.unregisteredClientsChan:

			// delete client from this map
			delete(cr.clientsMap, client)
			fmt.Println("chatroom: client unreg:", len(cr.clientsMap))

			// iterate and send message back
			for client := range cr.clientsMap {

				client.WsConn.WriteJSON(message.Message{

					Body: "User Disconnected...",
				})
			}
			break

		// case3. pop the message that's just received
		case msgObj := <-cr.messagesChan:
			fmt.Println("chatroom: pop and broadcast:", msgObj)

			// 0. upload all messages To S3 if redis is full
			cr.cache.UploadMessagesToS3(msgObj.RoomName)

			// 1. save to redis
			cr.cache.SaveMessageToStore(msgObj)

			for client := range cr.clientsMap {

				// 1. sends json message
				err := client.WsConn.WriteJSON(msgObj)

				// 2. handle error
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			break

		}

	}
}
