package model

import (
	"fmt"
	"strconv"
)

// Teammate represents a teammate with name, position in the team and his status
type Teammate struct {
	Name     string         `json:"name"`
	Position int            `json:"position"`
	Status   TeammateStatus `json:"status"`
}

// TeammateStatus represents the overall availability of a teammate
type TeammateStatus int

// possible values for TeammateStatus
const (
	StatusAvail TeammateStatus = iota
	StatusUnavail
	StatusSpare
)

// Key builds the db key for the teammate
func (m Teammate) Key() string {
	return fmt.Sprintf("mate:%v", m.ID())
}

// ID returns the short identifier for the teammate
func (m Teammate) ID() string {
	return fmt.Sprintf("%d", m.Position)
}

// FromID generates a stub object based on the given id. The stub object can then be used to get the full object from the repository
func (m *Teammate) FromID(id string) (err error) {
	m.Position, err = strconv.Atoi(id)
	return err
}
