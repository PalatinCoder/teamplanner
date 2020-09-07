package model

import (
	"encoding/json"
	"fmt"
)

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

// MarshalJSON is a custom json marshaller for votes. It replaces the embedded structs with their IDs as references
func (m Vote) MarshalJSON() ([]byte, error) {
	type Alias Vote
	return json.Marshal(&struct {
		Teammate string `json:"teammate"`
		Match    string `json:"match"`
		*Alias
	}{
		Teammate: m.Teammate.ID(),
		Match:    m.Match.ID(),
		Alias:    (*Alias)(&m),
	})
}

// UnmarshalJSON is a custom json unmarshaller for votes. It generates stubs of the embedded structs based on the IDs. Note that the stubs need to be fully inflated afterwards.
func (m *Vote) UnmarshalJSON(data []byte) error {
	type Alias Vote
	aux := &struct {
		Teammate string `json:"teammate"`
		Match    string `json:"match"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.Teammate.FromID(aux.Teammate)
	m.Match.FromID(aux.Match)
	return nil
}
