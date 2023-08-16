package message

type Message struct {
	// when it's converted to JSON, it will have the key name type.
	Type     int    `json:"type"` // type 1 means text frame
	Body     string `json:"body"`
	UserName string `json:"userName"`
}
