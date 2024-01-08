package services

import (
	"TChat/domain"
	"TChat/repositories"
	"time"
)

type message struct {
	messageRepository repositories.MessageRepository
}

func (m *message) CreateMessage(chatMessage domain.ChatMessage) error {
	return m.messageRepository.CreateMessage(chatMessage)
}

func (m *message) ListMessage(sendID string, receiveID string, startTime time.Time, endTime time.Time) ([]domain.ChatMessage, error) {
	msg, err := m.messageRepository.ListMessage(sendID, receiveID, startTime, endTime)
	if err != nil {
		return nil, err
	}
	var msgs []domain.ChatMessage
	for i := 0; i < len(msg); i++ {
		msgs = append(msgs, domain.ChatMessage{
			MessageID: msg[i].MessageID,
			SendID:    msg[i].SendID,
			ReceiveID: msg[i].ReceiveID,
			Message:   msg[i].Message,
			CreatedAt: msg[i].CreatedAt,
		})
	} // convert msg(models) to msgs(domain)
	return msgs, nil
}

func (m *message) DeleteMessage(messageID string) error {
	return m.messageRepository.DeleteMessage(messageID)
}

type MessageService interface {
	CreateMessage(chatMessage domain.ChatMessage) error
	ListMessage(sendID string, receiveID string, startTime time.Time, endTime time.Time) ([]domain.ChatMessage, error)
	DeleteMessage(messageID string) error
}

func NewMessageService(messageRepository repositories.MessageRepository) MessageService {
	return &message{
		messageRepository: messageRepository,
	}
}
