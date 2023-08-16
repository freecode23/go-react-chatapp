package socketUtil

import (
	"encoding/json"
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
		messageType, jsonBytes, err := c.wsConn.ReadMessage()
		if err != nil {
			log.Println(err)
			return

		}

		// 2. init var to store the message
		var msg message.Message

		// 3. handle error
		if err := json.Unmarshal(jsonBytes, &msg); err != nil {
			log.Println("Failed to unmarshal:", err)
			continue // Skip this iteration and go to next
		}

		// 4. assign type
		msg.Type = messageType

		// 5. write on channel
		c.chatroom.messagesChan <- msg

		fmt.Println("\nclient: push msg username:", msg.UserName, "body:", msg.Body)
	}

}
