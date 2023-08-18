package restUtil

import (
	"encoding/json"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
)

type apiHandler struct {
	Cache cache.Cache
}

func (a *apiHandler) getChatHistory(w http.ResponseWriter, r *http.Request) {

	// 1. Fetch the last 30 messages
	messagesStruct, err := a.Cache.GetLastMessagesStruct()

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
