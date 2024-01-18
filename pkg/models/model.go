package models

import "time"

type Message struct {
	MessageID string    `gorm:"column:message_id"`
	SendID    string    `gorm:"column:send_id"`
	ReceiveID string    `gorm:"column:receive_id"`
	Message   string    `gorm:"column:message"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type GenderType int

const (
	Male GenderType = iota
	Fale
	Other
)

type User struct {
	UserID      string     `gorm:"column:user_id"`
	PhoneNumber string     `gorm:"column:phone_number"`
	DateOfBirth time.Time  `gorm:"column:date_of_birth"`
	Name        string     `gorm:"column:name"`
	Email       string     `gorm:"column:email"`
	Gender      GenderType `gorm:"column:gender"`
	UserName    string     `gorm:"column:user_name"`
	Password    string     `gorm:"column:password"`
	AvtImg      string     `gorm:"column:avt_img"`
	CoverImg    string     `gorm:"column:cover_img"`
}

type Friend struct {
	UserID1 string `gorm:"column:user_id_1"`
	UserID2 string `gorm:"column:user_id_2"`
}

type Group struct {
	GroupID string `gorm:"group_id"`
	Name    string `gorm:"name"`
	Avt     string `gorm:"avt"`
}

type GroupMember struct {
	GroupID string `gorm:"group_id"`
	UserID  string `gorm:"user_id"`
}
