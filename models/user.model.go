package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	FirstName string
	LastName  string
	Username  string `gorm:"uniqueIndex;not null"`
	Password  string `json:"-"`

	Otp_enabled  bool `json:"-"`
	Otp_verified bool `json:"-"`

	Otp_secret   string `json:"-"`
	Otp_auth_url string `json:"-"`
}

func (user *User) BeforeCreate(*gorm.DB) error {
	user.ID = uuid.NewV4()

	return nil
}

type UserResponseStructure struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

type OTPInput struct {
	UserId string `json:"user_id"`
	Token  string `json:"otp"`
}
