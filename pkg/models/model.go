package models

import "time"

type Message struct {
	MessageID string    `gorm:"message_id"`
	SendID    string    `gorm:"send_id"`
	ReceiveID string    `gorm:"receive_id"`
	Message   string    `gorm:"message"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type GenderType int

const (
	Male GenderType = iota
	Fale
	Other
)

type User struct {
	UserID      string     `gorm:"user_id"`
	PhoneNumber string     `gorm:"phone_number"`
	DateOfBirth time.Time  `gorm:"date_of_birth"`
	Name        string     `gorm:"name"`
	Email       string     `gorm:"email"`
	Gender      GenderType `gorm:"gender"`
	UserName    string     `gorm:"user_name"`
	Password    string     `gorm:"password"`
	AvtImg      string     `gorm:"avt_img"`
	CoverImg    string     `gorm:"cover_img"`
}
