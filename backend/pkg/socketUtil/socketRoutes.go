package socketUtil

import (
	"fmt"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/freecode23/go-react-chatapp/pkg/chat"
	"github.com/gorilla/mux"
)

/*
*
 1. 0 Create or get chatroom for this client
    1 init go routine thread
    2 Create websocketconnection
    3 Create the client and add to registered client channel
    4 listening to this client connection

*
*/
func handleNewClient(cm *chat.ChatroomManager, cacheIf cache.Cache, w http.ResponseWriter, r *http.Request) {

	// 0. Create or get chatroom for this client
	// - Extract chatroom name from the URL
	vars := mux.Vars(r)
	chatroomName, exists := vars["chatroomName"]
	if !exists {
		http.Error(w, "Chatroom not specified", http.StatusBadRequest)
		return
	}

	// - check if chatroom already exist in map, if not create a new one and add to map in manager
	chatroomPtr := cm.GetOrCreateChatroom(cacheIf, chatroomName)

	// 1. go routine thread
	// set up a single chatroom in the background
	// will run concurrently without blocking the main thread.
	go chatroomPtr.ProcessChatroomEvents()

	// 2. upgrade http connection to web socket
	conn, err := upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// 3. create a client pointer
	clientPtr := &chat.Client{
		WsConn:   conn,
		Chatroom: chatroomPtr,
	}

	// - add clientPtr to the chatroom's Register
	chatroomPtr.RegisteredClientsChan <- clientPtr

	// 4. read messages from client and broadcast them (infinite loop)
	clientPtr.ListenMessages()

}

/*
*
Interface can take in a pointer or an object
we want to change the data in chatroom manager, so we should pass a pointer
*
*/
func SetupWebsocketRoutes(cacheIf cache.Cache, cm *chat.ChatroomManager) {

	// 2. Define the callback function when client calls
	websocketHandlerCallback := func(w http.ResponseWriter, r *http.Request) {
		handleNewClient(cm, cacheIf, w, r)
	}

	// 2. set up what to do on this route
	mRouter := mux.NewRouter()
	mRouter.HandleFunc("/start/{chatroomName}", websocketHandlerCallback)

	// 3. serve
	fmt.Println("Socketapi: Starting port 9000")
	err := http.ListenAndServe(":9000", mRouter)
	if err != nil {
		fmt.Println("Error starting socket:", err)
	}
}
