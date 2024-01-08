package domain

import "time"

type ChatMessage struct {
	MessageID string    `json:"message_id"`
	SendID    string    `json:"send_id"`
	ReceiveID string    `json:"receive_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
