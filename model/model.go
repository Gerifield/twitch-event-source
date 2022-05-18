package model

// MessageType .
type MessageType string

var (
	// ChatMessage .
	ChatMessage MessageType = "chat_message"

	// UserNotice .
	UserNotice MessageType = "user_notice"

	// ChannelPointRedeem .
	ChannelPointRedeem MessageType = "channel_point_redeem"
)

// EventFrame .
type EventFrame struct {
	ChannelID string      `json:"channel_id"`
	Type      MessageType `json:"type"`
	Payload   interface{} `json:"payload"`
}
