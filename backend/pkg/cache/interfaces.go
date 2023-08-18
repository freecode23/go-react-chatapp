package cache

import "github.com/freecode23/go-react-chatapp/pkg/message"

type Cache interface {
	GetLastMessagesStruct(roomName string) ([]message.Message, error)
	SaveMessageToStore(message message.Message) error
	UploadMessagesToS3(roomName string)
}
