package chat

import (
	"fmt"

	"github.com/freecode23/go-react-chatapp/pkg/cache"
)

type ChatroomManager struct {
	// This might contain a map of all chat rooms, methods to create or delete chat rooms, etc.
	chatroomsMap map[string]*Chatroom
}

func NewChatroomManager() *ChatroomManager {
	return &ChatroomManager{
		chatroomsMap: make(map[string]*Chatroom),
	}
}

// use pointer as receiver since we want to change the property of the chatroom manager
// always return a pointer
func (cm *ChatroomManager) GetOrCreateChatroom(cacheObj cache.Cache, chatroomName string) *Chatroom {
	// 1. return existing chatroom if exist
	if chatroom, exists := cm.chatroomsMap[chatroomName]; exists {
		return chatroom
	}
	// 2. create new chatroom
	newChatroomPtr := NewChatroom(cacheObj, chatroomName)

	// 3. add chat room to map
	cm.chatroomsMap[chatroomName] = newChatroomPtr
	fmt.Println("chatroomManager: currentRooms", cm.chatroomsMap)

	// 4. return new chatroom otherwise
	return newChatroomPtr
}

func (cm *ChatroomManager) GetAllChatroomsNames() []string {

	var chatroomNames []string

	for chatroomObject := range cm.chatroomsMap {
		chatroomNames = append(chatroomNames, chatroomObject)
	}
	return chatroomNames
}
