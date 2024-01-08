package domain

import "time"

type GenderType int

const (
	Male GenderType = iota
	Female
	Other
)

type User struct {
	UserID      string     `json:"user_id"`
	PhoneNumber string     `json:"phone_number"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Gender      GenderType `json:"gender"`
	UserName    string     `json:"user_name"`
	Password    string     `json:"password"`
	AvtImg      string     `json:"avt_img"`
	CoverImg    string     `json:"cover_img"`
}
