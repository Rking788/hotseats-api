package model

import (
	"fmt"
	"time"
)

// Event is a representation of an event that occurred at a game.
// Could be either a foulball or a homerun.
type Event struct {
	Eid       int       `gorm:"primary_key"`
	Type      string    `json:"type" binding:"required"`
	EventDate time.Time `json:"date" binding:"required"`
	Sid       int
	Stadium   Stadium
}

func (evt Event) String() string {
	return fmt.Sprintf("<Event: %p> Date=\"%s\", Type=\"%s\"", &evt, evt.EventDate, evt.Type)
}
