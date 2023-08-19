package main

import (
	"github.com/freecode23/go-react-chatapp/pkg/cache"
	"github.com/freecode23/go-react-chatapp/pkg/chat"
	"github.com/freecode23/go-react-chatapp/pkg/restUtil"
	"github.com/freecode23/go-react-chatapp/pkg/socketUtil"
)

func main() {
	// 0. init chatroom manager
	chatroomManagerPtr := chat.NewChatroomManager()

	// 1. init redis
	redisCachePtr := cache.NewRedisStore()

	// 2. init socket and rest API
	go socketUtil.SetupWebsocketRoutes(redisCachePtr, chatroomManagerPtr)
	restUtil.SetupRestRoutes(redisCachePtr, chatroomManagerPtr)

}
