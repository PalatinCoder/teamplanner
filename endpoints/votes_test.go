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
			testEndpoint(t, "getVotes()", e.getVotes, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}
