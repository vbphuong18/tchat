package dto

import "time"

type CreateMessageResponse struct {
	StatusError
}

type ListMessageResponse struct {
	StatusError
	Data []Message `json:"data"`
}

type DeleteMessageResponse struct {
	StatusError
}

type CreateMessageRequest struct {
	SendID    string `json:"send_id" validate:"required"`
	ReceiveID string `json:"receive_id" validate:"required"`
	Message   string `json:"message" validate:"min=1"`
}

type Message struct {
	MessageID string    `json:"message_id"`
	SendID    string    `json:"send_id"`
	ReceiveID string    `json:"receive_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
