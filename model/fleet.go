package model

import "github.com/google/uuid"

type Fleet struct {
	ID              uuid.UUID `gorm:"primaryKey" json:"id"`
	Plate           string    `json:"plate"`
	CurrentCapacity int       `json:"current_capacity"`
	MaxCapacity     int       `json:"max_capacity"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
}
