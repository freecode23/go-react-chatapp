package socketUtil

import (
	"log"
	"sync"

	"github.com/freecode23/go-react-chatapp/pkg/message"
	"github.com/gorilla/websocket"
)

type client struct {
	ID       string
	wsConn   *websocket.Conn
	chatroom *Chatroom
	mu       *sync.Mutex
}

/*
*
Get client's message and broadcast to chatroom's broadcast channel
*
*/
func (c *client) listenMessages() {

	//1.  will unregister if theres error or if client exit
	defer func() {
		c.chatroom.unregisteredClientsChan <- c
		c.wsConn.Close()
	}()

	// 2. infinite loop keep reading until client quits
	for {

		// 1. get the json message
		_, jsonBytes, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return

		}

		// 2. convert msg JsonBYtes to msgStruct Pointer
		msgStructPtr, err := message.NewMessageFromJSON(jsonBytes)
		if err != nil {
			log.Println("Failed to create message:", err)
			continue
		}

		// 5. write on channel
		c.chatroom.messagesChan <- *msgStructPtr
		// fmt.Printf("\nclient: push msgStructPtr:%v\n", msgStructPtr)
	}

}
