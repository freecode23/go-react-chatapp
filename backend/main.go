package main

import (
	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/freecode23/go-react-chatapp/pkg/restUtil"
	"github.com/freecode23/go-react-chatapp/pkg/socketUtil"
)

func main() {

	// 1. init redis
	redisCachePtr := cache.NewRedisStore()

	// 2. init socket and rest API
	go socketUtil.SetupWebsocketRoutes(redisCachePtr)
	restUtil.SetupRestRoutes(redisCachePtr)

}
