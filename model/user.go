package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Email       string    `gorm:"unique" json:"email"`
	Password    string    `json:"password"`
	KiriBalance float64   `json:"kiri_balance"`
	KiriPoint   float64   `json:"kiri_point"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `gorm:"notNull" json:"email"`
	Password string `gorm:"notNull" json:"password"`
}

type UserRegister struct {
	Name            string `gorm:"notNull" json:"name"`
	Email           string `gorm:"notNull" json:"email"`
	Password        string `gorm:"notNull" json:"password"`
	ConfirmPassword string `gorm:"notNull" json:"confirm_password"`
}
