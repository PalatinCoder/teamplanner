package endpoints

import (
	"io"
	"strings"
	"testing"

	"jan-sl.de/teamplanner/model"
)

func TestEndpoints_getTeammates(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives empty response", &MockRepository{mates: []model.Teammate{}}, nil, 200, "[]"},
		{"request body is ignored", &MockRepository{mates: []model.Teammate{}}, strings.NewReader("hello world"), 200, "[]"},
		{"gives back JSON", &MockRepository{mates: Teammates}, nil, 200, string(marshallJSONWithoutError(Teammates))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getTeammates()", e.getTeammates, nil, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}

func TestEndpoints_getTeammate(t *testing.T) {
	tests := []struct {
		name           string
		repository     model.Dataprovider
		vars           map[string]string
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{"empty repository gives 404 response", &MockRepository{mates: []model.Teammate{}}, map[string]string{"id": "1"}, nil, 404, `{"error":"not found"}`},
		{"gives back JSON", &MockRepository{mates: Teammates}, map[string]string{"id": Teammates[0].ID()}, nil, 200, string(marshallJSONWithoutError(Teammates[0]))},
		{"gives internal server error when repository goes wrong", &MockErrorRepository{}, map[string]string{"id": "1"}, nil, 500, `{"error":"mock error"}`},
	}

	for _, tt := range tests {
		e := &Endpoints{tt.repository, nil}

		t.Run(tt.name, func(t *testing.T) {
			testEndpoint(t, "getTeammate()", e.getTeammate, tt.vars, tt.body, tt.expectedStatus, tt.expectedBody)
		})
	}
}
