package endpoints

import (
	"errors"
	"time"

	. "jan-sl.de/teamplanner/model"
)

var Teammates = []Teammate{
	{Name: "Abcd", Position: 1, Status: StatusAvail},
	{Name: "Efgh", Position: 2, Status: StatusUnavail},
	{Name: "Jklr", Position: 3, Status: StatusSpare},
}

var Matches = []Match{
	{Date: time.Now().Truncate(12*time.Hour).AddDate(0, 0, -7), Description: "Letzte Woche"},
	{Date: time.Now().Truncate(12*time.Hour).AddDate(0, 0, 0), Description: "Heute"},
	{Date: time.Now().Truncate(12*time.Hour).AddDate(0, 0, 1), Description: "Morgen"},
	{Date: time.Now().Truncate(12*time.Hour).AddDate(0, 0, 7), Description: "Nächste Woche"},
}

var Votes = []Vote{
	{Teammate: Teammates[0], Match: Matches[0], Vote: VoteYes},
	{Teammate: Teammates[1], Match: Matches[0], Vote: VoteMaybe},
	{Teammate: Teammates[2], Match: Matches[0], Vote: VoteYes},
	{Teammate: Teammates[0], Match: Matches[1], Vote: VoteNo},
	{Teammate: Teammates[1], Match: Matches[1], Vote: VoteYes},
	{Teammate: Teammates[2], Match: Matches[1], Vote: VoteMaybe},
	{Teammate: Teammates[0], Match: Matches[2], Vote: VoteYes},
	{Teammate: Teammates[1], Match: Matches[2], Vote: VoteNo},
	{Teammate: Teammates[2], Match: Matches[2], Vote: VoteYes},
	{Teammate: Teammates[0], Match: Matches[3], Vote: VoteMaybe},
	{Teammate: Teammates[1], Match: Matches[3], Vote: VoteNo},
	{Teammate: Teammates[2], Match: Matches[3], Vote: VoteYes},
}

type MockRepository struct {
	matches []Match
	mates   []Teammate
	votes   []Vote
}

func (m *MockRepository) GetMatches() ([]Match, error) {
	return m.matches, nil
}
func (m *MockRepository) GetTeammates() ([]Teammate, error) {
	return m.mates, nil
}
func (m *MockRepository) GetVotes() ([]Vote, error) {
	return m.votes, nil
}
func (m *MockRepository) GetVotesByTeammate(Teammate) ([]Vote, error) {
	return m.votes, nil
}
func (m *MockRepository) GetVotesForMatch(Match) ([]Vote, error) {
	return m.votes, nil
}

//MockErrorRepository returns an error on every method
type MockErrorRepository struct{}

func (m *MockErrorRepository) GetMatches() ([]Match, error) {
	return nil, errors.New("mock error")
}
func (m *MockErrorRepository) GetTeammates() ([]Teammate, error) {
	return nil, errors.New("mock error")
}
func (m *MockErrorRepository) GetVotes() ([]Vote, error) {
	return nil, errors.New("mock error")
}
func (m *MockErrorRepository) GetVotesByTeammate(Teammate) ([]Vote, error) {
	return nil, errors.New("mock error")
}
func (m *MockErrorRepository) GetVotesForMatch(Match) ([]Vote, error) {
	return nil, errors.New("mock error")
}
