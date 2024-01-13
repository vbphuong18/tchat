package repositories

import (
	"TChat/pkg/domain"
	"TChat/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type message struct {
	db *gorm.DB
}

func (m *message) CreateMessage(chatMessage domain.ChatMessage) error {
	return m.db.Create(&models.Message{
		MessageID: uuid.New().String(),
		SendID:    chatMessage.SendID,
		ReceiveID: chatMessage.ReceiveID,
		Message:   chatMessage.Message,
		CreatedAt: chatMessage.CreatedAt,
		UpdatedAt: chatMessage.CreatedAt,
	}).Error
}

func (m *message) ListMessage(sendID string, receiveID string, startTime time.Time, endTime time.Time) ([]models.Message, error) {
	var msg []models.Message
	err := m.db.Where("((send_id = ? and receive_id = ?) or (receive_id = ? and send_id = ?)) and created_at > ? and created_at < ?", sendID, receiveID, sendID, receiveID, startTime, endTime).Find(&msg).Error
	return msg, err
}

func (m *message) DeleteMessage(messageID string) error {
	return m.db.Where("message_id = ?", messageID).Delete(&models.Message{}).Error
}

type MessageRepository interface {
	CreateMessage(chatMessage domain.ChatMessage) error
	ListMessage(sendID string, receiveID string, startTime time.Time, endTime time.Time) ([]models.Message, error)
	DeleteMessage(messageID string) error
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &message{
		db: db,
	}
}
