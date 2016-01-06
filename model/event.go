package model

// Event is a representation of an event that occurred at a game.
// Could be either a foulball or a homerun.
type Event struct {
	EventType string `json:"type"`
	Date      string `json:"date"`
}
