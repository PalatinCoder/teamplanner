package model

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/buntdb"
)

// Dataprovider is the general data repository
type Dataprovider interface {
	//GetMatches returns all matches
	GetMatches() ([]Match, error)
	// GetMatch inflates a match into the given object.
	// The object must have all required properties set for a call to Model.Key() to succeed
	GetMatch(*Match) error
	//GetTeammates returns all teammates
	GetTeammates() ([]Teammate, error)
	// GetTeammate inflates a teammate into the given object.
	// The object must have all required properties set for a call to Model.Key() to succeed
	GetTeammate(*Teammate) error
	GetVotes() ([]Vote, error)
	GetVotesByTeammate(mate Teammate) ([]Vote, error)
	GetVotesForMatch(match Match) ([]Vote, error)
}

// BuntDb implements Dataprovider with buntdb as K/V store
type BuntDb struct {
	db *buntdb.DB
}

// NewBuntDb initializes a new repository
func NewBuntDb(db *buntdb.DB) *BuntDb {
	r := &BuntDb{db}
	createIndexes(db)
	return r
}

func createIndexes(db *buntdb.DB) {
	db.CreateIndex("teammates", "mate:*", buntdb.IndexJSON("position"))
	db.CreateIndex("matches", "match:*", buntdb.IndexJSON("date"))
	db.CreateIndex("votes", "vote:*", buntdb.IndexJSON("match.date"))
}

// GetTeammates retrieves all teammates
func (r *BuntDb) GetTeammates() ([]Teammate, error) {
	var mates []Teammate

	err := r.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("teammates", func(key, value string) bool {
			var m Teammate
			json.Unmarshal([]byte(value), &m)
			mates = append(mates, m)
			return true
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return mates, nil
}

// GetTeammate retrieves a single teammate. The given Teammate object must have all attributes set so that a call to Teammate.Key() will be successfull.
// The given object will then be inflated with the remaining properties.
func (r *BuntDb) GetTeammate(mate *Teammate) error {
	err := r.db.View(func(tx *buntdb.Tx) error {
		j, err := tx.Get(mate.Key(), false)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(j), &mate)
		return err
	})
	return err
}

// GetMatches retrieves all match dates
func (r *BuntDb) GetMatches() ([]Match, error) {
	var matches []Match

	err := r.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("matches", func(key, value string) bool {
			var m Match
			json.Unmarshal([]byte(value), &m)
			matches = append(matches, m)
			return true
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return matches, nil
}

// GetMatch inflates a match date
func (r *BuntDb) GetMatch(match *Match) error {
	err := r.db.View(func(tx *buntdb.Tx) error {
		j, err := tx.Get(match.Key(), false)
		if err != nil {
			return err
		}
		err = json.Unmarshal([]byte(j), &match)
		return err
	})
	return err
}

// GetVotes retrieves all votes
func (r *BuntDb) GetVotes() ([]Vote, error) {
	var votes []Vote

	err := r.db.View(func(tx *buntdb.Tx) error {
		err := tx.Ascend("votes", func(key, value string) bool {
			var m Vote
			json.Unmarshal([]byte(value), &m)
			votes = append(votes, m)
			return true
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return votes, nil
}

// GetVotesByTeammate retrieves all votes by a single teammate
func (r *BuntDb) GetVotesByTeammate(mate Teammate) ([]Vote, error) {
	var votes []Vote
	pattern := fmt.Sprintf("vote:%s:*", mate.ID())

	err := r.db.View(func(tx *buntdb.Tx) error {
		err := tx.AscendKeys(pattern, func(key, value string) bool {
			var m Vote
			json.Unmarshal([]byte(value), &m)
			votes = append(votes, m)
			return true
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return votes, nil
}

// GetVotesForMatch retrieves all votes for a matchdate
func (r *BuntDb) GetVotesForMatch(match Match) ([]Vote, error) {
	var votes []Vote
	pattern := fmt.Sprintf("vote:*:%s", match.ID())

	err := r.db.View(func(tx *buntdb.Tx) error {
		err := tx.AscendKeys(pattern, func(key, value string) bool {
			var m Vote
			json.Unmarshal([]byte(value), &m)
			votes = append(votes, m)
			return true
		})
		return err
	})

	if err != nil {
		return nil, err
	}
	return votes, nil
}
