package socketUtil

import (
	"fmt"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
)

/*
*
websocket chatroomis a struct from socketUtil package
*
*/
func handleNewClient(cr *chatroom, w http.ResponseWriter, r *http.Request) {

	// 1. upgrade http connection to web socket
	conn, err := upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// 2. create a client pointer
	clientPtr := &client{
		wsConn:   conn,
		chatroom: cr,
	}

	// 3. add clientPtr to the chatroom's Register
	cr.registeredClientsChan <- clientPtr

	// 4. read messages from client and broadcast them (infinite loop)
	clientPtr.listenMessages()
}

func SetupWebsocketRoutes(cacheIf cache.Cache) {

	// 1. init chatroom/ chatroom
	chatroomPtr := newChatroom(cacheIf)

	// 2. go routine thread
	// set up a single chatroom in the background
	// will run concurrently without blocking the main thread.
	go chatroomPtr.processChatroomEvents()

	// 3. Define the callback function when client calls
	websocketHandlerCallback := func(w http.ResponseWriter, r *http.Request) {
		handleNewClient(chatroomPtr, w, r)
	}

	// 4. set up what to do on this route
	http.HandleFunc("/start", websocketHandlerCallback)

	// 5. serve
	fmt.Println("Socketapi: Starting port 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Error starting socket:", err)
	}
}
