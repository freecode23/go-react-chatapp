package main

import (
	"fmt"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/wsutil"
)

/*
*
websocket Chatroom is a struct from wsutil package
*
*/
func handleNewClient(chatroom *wsutil.Chatroom, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached. serving")

	// 1. upgrade http connection to web socket
	conn, err := wsutil.Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// 2. create a client pointer
	clientPtr := &wsutil.Client{
		Conn:     conn,
		Chatroom: chatroom,
	}

	// 3. add clientPtr to the chatroom's Register
	chatroom.RegisteredClientsChan <- clientPtr

	// 4. read messages from client and broadcast them (infinite loop)
	clientPtr.ListenMessages()
}

func setupRoutes() {

	// 1. init Chatroom / chatroom
	chatroomPtr := wsutil.NewChatroom()

	// 2. go routine thread
	// set up a single chatroom in the background
	// will run concurrently without blocking the main thread.
	go chatroomPtr.ProcessChatroomEvents()

	// 3. Define the callback function when client calls
	websocketHandlerCallback := func(w http.ResponseWriter, r *http.Request) {
		handleNewClient(chatroomPtr, w, r)
	}

	// 4. set up what to do on this route
	http.HandleFunc("/start", websocketHandlerCallback)
}

func main() {
	fmt.Println("Chat App Backend")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
