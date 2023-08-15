package main

import (
	"fmt"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/freecode23/go-react-chatapp/pkg/restUtil"
	"github.com/freecode23/go-react-chatapp/pkg/socketUtil"
)

func main() {
	fmt.Println("Chat App Backend")

	// 2. init redis
	redisCachePtr := cache.NewRedisStore()

	// 3. init socket and rest API
	go socketUtil.SetupWebsocketRoutes(redisCachePtr)
	restUtil.SetupRestRoutes(redisCachePtr)

}
