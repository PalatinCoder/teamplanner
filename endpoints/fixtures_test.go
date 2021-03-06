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
	{Date: time.Now().Truncate(24*time.Hour).AddDate(0, 0, -7), Description: "Letzte Woche"},
	{Date: time.Now().Truncate(24*time.Hour).AddDate(0, 0, 0), Description: "Heute"},
	{Date: time.Now().Truncate(24*time.Hour).AddDate(0, 0, 1), Description: "Morgen"},
	{Date: time.Now().Truncate(24*time.Hour).AddDate(0, 0, 7), Description: "Nächste Woche"},
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

// filterVotes calls f() on every vote in the slice to determine if it should be included in the filtered slice
func filterVotes(vs []Vote, f func(Vote) bool) []Vote {
	vsf := make([]Vote, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
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
func (m *MockRepository) GetTeammate(mate *Teammate) error {
	if len(m.mates) == 0 {
		return errors.New("not found")
	}
	*mate = m.mates[0]
	return nil
}
func (m *MockRepository) SetTeammate(mate *Teammate) (bool, error) {
	return false, nil
}
func (m *MockRepository) GetMatch(match *Match) error {
	if len(m.matches) == 0 {
		return errors.New("not found")
	}
	*match = m.matches[1]
	return nil
}
func (m *MockRepository) SetMatch(match *Match) (bool, error) {
	return false, nil
}
func (m *MockRepository) GetVotes() ([]Vote, error) {
	return m.votes, nil
}
func (m *MockRepository) GetVotesByTeammate(mate Teammate) ([]Vote, error) {
	return filterVotes(m.votes, func(v Vote) bool { return v.Teammate.ID() == mate.ID() }), nil
}
func (m *MockRepository) GetVotesForMatch(match Match) ([]Vote, error) {
	return filterVotes(m.votes, func(v Vote) bool { return v.Match.ID() == match.ID() }), nil
}
func (m *MockRepository) SetVote(vote *Vote) (bool, error) {
	return false, nil
}

//MockErrorRepository returns an error on every method
type MockErrorRepository struct{}

func (m *MockErrorRepository) GetMatches() ([]Match, error) {
	return nil, errors.New("mock error")
}
func (m *MockErrorRepository) GetTeammates() ([]Teammate, error) {
	return nil, errors.New("mock error")
}
func (m *MockErrorRepository) GetTeammate(mate *Teammate) error {
	return errors.New("mock error")
}
func (m *MockErrorRepository) SetTeammate(*Teammate) (bool, error) {
	return false, errors.New("mock error")
}
func (m *MockErrorRepository) GetMatch(match *Match) error {
	return errors.New("mock error")
}
func (m *MockErrorRepository) SetMatch(*Match) (bool, error) {
	return false, errors.New("mock error")
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
func (m *MockErrorRepository) SetVote(*Vote) (bool, error) {
	return false, errors.New("mock error")
}
