package cache

import "github.com/freecode23/go-react-chatapp/pkg/message"

type Cache interface {
	GetLast30Messages() ([]message.Message, error)
	SaveMessageToStore(message message.Message) error
}
