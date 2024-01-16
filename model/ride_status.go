package model

import (
	"github.com/google/uuid"
	"time"
)

type Status struct {
	OrderID   string    `gorm:"primaryKey" json:"order_id"`
	User      User      `gorm:"ForeignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uuid.UUID `gorm:"null" json:"user_id"`
	Status    bool      `json:"status"`
	Start     time.Time `json:"start"`
	End       time.Time `json:"end"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
