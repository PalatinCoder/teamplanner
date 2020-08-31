package model

import "fmt"

// VoteOption is a teammate's availiability for a specific match date
type VoteOption int

// possible values for VoteOption
const (
	VoteYes VoteOption = iota
	VoteNo
	VoteMaybe
)

// Vote represents a vote given by a teammate on a specific match date
type Vote struct {
	Teammate Teammate   `json:"teammate,omitempty"`
	Match    Match      `json:"match,omitempty"`
	Vote     VoteOption `json:"vote"`
}

// Key builds the db key for the teammate
func (m Vote) Key() string {
	return fmt.Sprintf("vote:%v", m.ID())
}

// ID returns the ID of the vote
func (m Vote) ID() string {
	return fmt.Sprintf("%s:%s", m.Teammate.ID(), m.Match.ID())
}
