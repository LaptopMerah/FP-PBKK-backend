package entity

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	EventName string       `json:"event_name"`
	Date      string       `json:"date"`
	Location  string       `json:"location"`
	Details   string       `json:"details"`
	Participants []Participant `gorm:"foreignKey:EventID" json:"participants"`
}
