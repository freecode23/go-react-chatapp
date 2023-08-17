package message

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	// when it's converted to JSON, it will have the key name type.
	ID        string    `json:"id"`
	UserName  string    `json:"userName"`
	RoomName  string    `json:"roomName"`
	Timestamp time.Time `json:"timestamp"`
	Body      string    `json:"body"`
}

func NewMessageFromJSON(jsonBytes []byte) (*Message, error) {
	var msgStruct Message
	if err := json.Unmarshal(jsonBytes, &msgStruct); err != nil {
		return nil, err
	}

	// create uuid without dash
	msgStruct.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	msgStruct.Timestamp = time.Now()

	return &msgStruct, nil
}
