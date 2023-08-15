package wsutil

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	Conn     *websocket.Conn
	Chatroom *Chatroom
	mu       *sync.Mutex
}

type Message struct {
	// when it's converted to JSON, it will have the key name type.
	Type int    `json:"type"` // type 1 means text frame
	Body string `json:"body"`
}

/*
*
Get client's message and broadcast to chatroom's broadcast channel
*
*/
func (c *Client) ListenMessages() {

	//1.  will only be executed if theres error
	// or if client exits
	defer func() {
		c.Chatroom.UnregisteredClientsChan <- c
		c.Conn.Close()
	}()

	// 2. infinite loop keep reading until client quits
	for {

		// 1. get the json message
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return

		}
		// 2. create the mssage
		message := Message{Type: messageType, Body: string(p)}

		// 3. write on channel
		c.Chatroom.MessagesChan <- message

		fmt.Printf("Message received: %+v\n", message)
	}

}
