package wsutil

import (
	"fmt"
)

type Chatroom struct {
	RegisteredClientsChan   chan *Client
	UnregisteredClientsChan chan *Client
	Clients                 map[*Client]bool // key client, value bool. made up set like in python
	MessagesChan            chan Message
}

/*
*
Init new Chatroom which is just a collection of clients that are connected together
Think of this as a single chat room
Return reference to a new Chatroom
*
*/
func NewChatroom() *Chatroom {

	return &Chatroom{
		RegisteredClientsChan:   make(chan *Client),
		UnregisteredClientsChan: make(chan *Client),
		Clients:                 make(map[*Client]bool),
		MessagesChan:            make(chan Message),
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
func (chatroom *Chatroom) ProcessChatroomEvents() {

	for {
		select {

		// case1: read the client that register
		// client here looks like this: memory address &{ 0x140001542c0 0x1400006e0e0 <nil>}
		// does not print out the actual values stored in pointers
		case client := <-chatroom.RegisteredClientsChan:

			// 1. insert client
			chatroom.Clients[client] = true
			fmt.Println("size of connection Chatroom:", len(chatroom.Clients))

			for client := range chatroom.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "New User Joined...",
				})
			}
			break

		// case2. read the client that unregister
		case client := <-chatroom.UnregisteredClientsChan:

			// delete client from this map
			delete(chatroom.Clients, client)
			fmt.Println("size of connection chatroom:", len(chatroom.Clients))

			// iterate and send message back
			for client := range chatroom.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{
					Type: 1,
					Body: "User Disconnected...",
				})
			}
			break

		// cas3. pop the message that's just received
		case msg := <-chatroom.MessagesChan:
			fmt.Println("Sending message to all clients in chatroom")

			for client := range chatroom.Clients {

				// 1. sends json message
				err := client.Conn.WriteJSON(msg)

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
