package restUtil

import (
	"encoding/json"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/gorilla/mux"
)

type apiHandler struct {
	Cache cache.Cache
}

func (a *apiHandler) getChatHistory(w http.ResponseWriter, r *http.Request) {
	// 0. Extract chatroom name from the URL
	vars := mux.Vars(r)
	chatroomName, exists := vars["chatroomName"]
	if !exists {
		http.Error(w, "Chatroom not specified", http.StatusBadRequest)
		return
	}

	// 1. Fetch the last messages
	messagesStruct, err := a.Cache.GetLastMessagesStruct(chatroomName)

	if err != nil {
		http.Error(w, "Failed to fetch chat history", http.StatusInternalServerError)
		return
	}

	// 2. Convert the messagesStruct to JSON bytes
	jsonBytes, err := json.Marshal(messagesStruct)
	if err != nil {
		http.Error(w, "Failed to convert chat history to JSON", http.StatusInternalServerError)
		return
	}

	// 3. Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")

	// 4. send back to client
	w.Write(jsonBytes)
}
