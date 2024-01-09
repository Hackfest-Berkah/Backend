package model

import (
	"github.com/google/uuid"
	"time"
)

type History struct {
	OrderID   string    `gorm:"unique" json:"order_id"`
	User      User      `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uuid.UUID `gorm:"null" json:"user_id"`
	Type      string    `json:"type"`
	Plate     string    `json:"plate"`
	Amount    string    `json:"amount"`
	Time      string    `json:"time"`
	CreatedAt time.Time `json:"created_at"`
}
