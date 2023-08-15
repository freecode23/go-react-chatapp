package main

import (
	"fmt"
	"net/http"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/freecode23/go-react-chatapp/pkg/restUtil"
	"github.com/freecode23/go-react-chatapp/pkg/socketUtil"
)

func main() {
	fmt.Println("Chat App Backend")

	// 2. init redis
	redisCachePtr := cache.NewRedisStore()

	// 3. init socket and rest API
	socketUtil.SetupWebsocketRoutes(redisCachePtr)
	restUtil.SetupRestRoutes(redisCachePtr)
	http.ListenAndServe(":9000", nil)
}
