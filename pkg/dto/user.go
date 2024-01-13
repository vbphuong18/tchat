package dto

import "time"

type GenderType int

const (
	Male GenderType = iota
	FeMale
	Other
)

type CreateUserResponse struct {
	StatusError
}

type ListUserResponse struct {
	StatusError
	Data []User `json:"data"`
}

type SearchUserResponse struct {
	StatusError
	Data []User `json:"data"`
}

type GetUserByUserIDResponse struct {
	StatusError
	Data User `json:"data"`
}

type DeleteUserResponse struct {
	StatusError
}

type CreateUserRequest struct {
	PhoneNumber string     `json:"phone_number" validate:"required"`
	DateOfBirth time.Time  `json:"date_of_birth" validate:"required"`
	Name        string     `json:"name" validate:"required"`
	Email       string     `json:"email" validate:"required,email"`
	Gender      GenderType `json:"gender" validate:"max=2"`
	UserName    string     `json:"user_name" validate:"required"`
	Password    string     `json:"password" validate:"required,min=8"`
	AvtImg      string     `json:"avt_img"`
	CoverImg    string     `json:"cover_img"`
}

type User struct {
	UserID      string     `json:"user_id"`
	PhoneNumber string     `json:"phone_number,omitempty"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	Name        string     `json:"name"`
	Email       string     `json:"email,omitempty"`
	Gender      GenderType `json:"gender"`
	UserName    string     `json:"user_name,omitempty"`
	Password    string     `json:"password,omitempty"`
	AvtImg      string     `json:"avt_img"`
	CoverImg    string     `json:"cover_img"`
}
