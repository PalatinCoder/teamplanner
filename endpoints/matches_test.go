package endpoints

import (
	"io"
	"strings"
	"testing"

	"jan-sl.de/teamplanner/model"
)

func TestEndpoints_getMatches(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives empty response", &MockRepository{matches: []model.Match{}}, nil, 200, "[]"},
		{"request body is ignored", &MockRepository{matches: []model.Match{}}, strings.NewReader("hello world"), 200, "[]"},
		{"gives back JSON", &MockRepository{matches: Matches}, nil, 200, string(marshallJSONWithoutError(Matches))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getMatches()", e.getMatches, nil, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}

func TestEndpoints_getMatch(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		vars           map[string]string
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives 404 response", &MockRepository{matches: []model.Match{}}, map[string]string{"id": "20060102"}, nil, 404, `{"error":"not found"}`},
		{"gives back JSON", &MockRepository{matches: Matches}, map[string]string{"id": Matches[1].ID()}, nil, 200, string(marshallJSONWithoutError(Matches[1]))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, map[string]string{"id": "20060102"}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getMatch()", e.getMatch, tt.vars, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}
