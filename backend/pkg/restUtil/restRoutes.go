package restUtil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/gorilla/mux"
)

type apiHandler struct {
	Cache cache.Cache
}

func (a *apiHandler) getChatHistory(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch the last 30 messages
	messages, err := a.Cache.GetLast30Messages()
	if err != nil {
		http.Error(w, "Failed to fetch chat history", http.StatusInternalServerError)
		return
	}

	// 2. Convert the messages to JSON
	jsonData, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "Failed to convert chat history to JSON", http.StatusInternalServerError)
		return
	}

	// 3. Set the Content-Type header and write the JSON response
	w.Header().Set("Content-Type", "application/json")

	// 4. send back to client
	w.Write(jsonData)
}

/*
*
The argument type is interface. but if we pass on redisCachePtr
it will hold a value of redisCachePtr and we can use it to
retrieve chatHIstory from the cache
*
*/
func SetupRestRoutes(cacheIf cache.Cache) {

	// 1. init muxRouter and apiHandler
	muxRouter := mux.NewRouter()
	ah := &apiHandler{
		Cache: cacheIf,
	}

	// 2. assign routes and method
	muxRouter.HandleFunc("/chatHistory", ah.getChatHistory).Methods("GET")

	// 3. make muxRouter handle requests to the root path
	http.Handle("/", muxRouter)

	fmt.Println("Restapi: Starting port 8080")
	http.ListenAndServe(":8080", nil)
}
