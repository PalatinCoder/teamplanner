package model

import (
	"fmt"
	"time"
)

// Match represents a match date along with a description
type Match struct {
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}

// Key returns the db key for the match
func (m Match) Key() string {
	return fmt.Sprintf("match:%s", m.ID())
}

// ID returns the short id for the match date
func (m Match) ID() string {
	return fmt.Sprintf("%s", m.Date.Format("20060102"))
}
