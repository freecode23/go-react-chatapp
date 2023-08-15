package socketUtil

import (
	"fmt"
	"log"
	"sync"

	"github.com/freecode23/go-react-chatapp/pkg/message"
	"github.com/gorilla/websocket"
)

type client struct {
	ID       string
	wsConn   *websocket.Conn
	chatroom *chatroom
	mu       *sync.Mutex
}

/*
*
Get client's message and broadcast to chatroom's broadcast channel
*
*/
func (c *client) listenMessages() {

	//1.  will only be executed if theres error
	// or if client exits
	defer func() {
		c.chatroom.unregisteredClientsChan <- c
		c.wsConn.Close()
	}()

	// 2. infinite loop keep reading until client quits
	for {

		// 1. get the json message
		messageType, p, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return

		}
		// 2. create the mssage
		message := message.Message{Type: messageType, Body: string(p)}

		// 3. write on channel
		c.chatroom.messagesChan <- message

		fmt.Printf("\nclient: push %+v\n", message)
	}

}
