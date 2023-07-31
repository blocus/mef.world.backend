package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Habit struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Title     string     `gorm:"type:varchar(255);not null"`
	StartDate *time.Time `gorm:"not null;autoCreateTime:milli"`
	Active    bool       `gorm:"type:boolean;not null;default:true;"`
	UserID    uuid.UUID  `json:"-"`
	User      User       `gorm:"foreignKey:UserID"`
}

type HabitActivity struct {
	ID      uuid.UUID  `gorm:"type:uuid;primary_key;"`
	Date    *time.Time `gorm:"not null;autoCreateTime:milli"`
	Done    bool       `gorm:"type:boolean;not null;default:false;"`
	HabitID string     `gorm:"type:uuid(255);not null"`
	Habit   Habit      `gorm:"foreignKey:HabitID"`
}

type ActivityResponse struct {
	ID   uuid.UUID  `json:"id"`
	Date *time.Time `json:"date"`
	Done bool       `json:"done"`
}

func (habit *Habit) BeforeCreate(*gorm.DB) error {
	habit.ID = uuid.NewV4()
	return nil
}

func (habitActivity *HabitActivity) BeforeCreate(*gorm.DB) error {
	habitActivity.ID = uuid.NewV4()
	return nil
}

type HabitResponseStructure struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	StartDate *time.Time `json:"start_at"`
	Active    bool       `json:"active"`
	User      User       `json:"user"`
}

type HabitActivityResponse struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	StartDate *time.Time `json:"start_at"`
	Active    bool       `json:"active"`
	Activity  []ActivityResponse
}
