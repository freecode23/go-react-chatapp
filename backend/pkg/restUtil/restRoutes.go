package restUtil

import (
	"fmt"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

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
	muxRouter.HandleFunc("/chatHistory/{chatroomName}", ah.getChatHistory).Methods("GET")

	// 3. Use the CORS middleware for your router
	handler := cors.Default().Handler(muxRouter)

	// 4. make muxRouter handle requests to the root path
	http.Handle("/", muxRouter)

	fmt.Println("Restapi: Starting port 8080")
	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
