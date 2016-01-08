package model

import (
	"fmt"
)

// Event is a representation of an event that occurred at a game.
// Could be either a foulball or a homerun.
type Event struct {
	EventType string `json:"type" binding:"required"`
	Date      string `json:"date" binding:"required"`
}

func (evt Event) String() string {
	return fmt.Sprintf("<Event: %p> Date=\"%s\", Type=\"%s\"", &evt, evt.Date, evt.EventType)
}
