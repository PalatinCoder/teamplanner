package endpoints

import (
	"io"
	"strings"
	"testing"

	"jan-sl.de/teamplanner/model"
)

func TestEndpoints_getVotes(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives empty response", &MockRepository{votes: []model.Vote{}}, nil, 200, "[]"},
		{"request body is ignored", &MockRepository{votes: []model.Vote{}}, strings.NewReader("hello world"), 200, "[]"},
		{"gives back JSON", &MockRepository{votes: Votes}, nil, 200, string(marshallJSONWithoutError(Votes))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getVotes()", e.getVotes, nil, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}

func TestEndpoints_getVotesForTeammate(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		vars           map[string]string
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives 404 response", &MockRepository{votes: []model.Vote{}}, map[string]string{"id": "1"}, nil, 404, `{"error":"not found"}`},
		{"gives back JSON", &MockRepository{votes: Votes, mates: Teammates}, map[string]string{"id": Teammates[0].ID()}, nil, 200, string(marshallJSONWithoutError(filterVotes(Votes, func(v model.Vote) bool { return v.Teammate.ID() == Teammates[0].ID() })))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, map[string]string{"id": "1"}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getVotesForTeammate()", e.getVotesForTeammate, tt.vars, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}

func TestEndpoints_getVotesForMatch(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		vars           map[string]string
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives 404 response", &MockRepository{votes: []model.Vote{}}, map[string]string{"id": "20060102"}, nil, 404, `{"error":"not found"}`},
		{"gives back JSON", &MockRepository{votes: Votes, matches: Matches}, map[string]string{"id": Matches[1].ID()}, nil, 200, string(marshallJSONWithoutError(filterVotes(Votes, func(v model.Vote) bool { return v.Match.ID() == Matches[1].ID() })))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, map[string]string{"id": "20060102"}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getVotesForMatch()", e.getVotesForMatch, tt.vars, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}
