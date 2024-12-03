package entity

import "gorm.io/gorm"

type Participant struct {
	gorm.Model
	EventID uint   `json:"event_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}