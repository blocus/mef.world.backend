package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	FirstName string    `gorm:"type:varchar(255);not null"`
	LastName  string    `gorm:"type:varchar(255);not null"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`

	Otp_enabled  bool `gorm:"default:false;"`
	Otp_verified bool `gorm:"default:false;"`

	Otp_secret   string
	Otp_auth_url string
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
