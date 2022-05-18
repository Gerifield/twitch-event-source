package model

// MessageType .
type MessageType string

var (
	// ChatMessage .
	ChatMessage MessageType = "chat_message"

	UserNotice MessageType = "user_notice"
)

// EventFrame .
type EventFrame struct {
	ChannelID string      `json:"channel_id"`
	Type      MessageType `json:"type"`
	Payload   interface{} `json:"payload"`
}
