package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tidwall/buntdb"
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

func fillDb(db *buntdb.DB) {
	err := db.Update(func(tx *buntdb.Tx) error {
		for _, m := range Teammates {
			if err := insert(tx, m); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func insert(tx *buntdb.Tx, m Model) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}
	_, _, err = tx.Set(m.Key(), string(j), nil)
	if err != nil {
		return err
	}
	return nil
}

// calls json.Marshall and panics on error
func marshallJSONWithoutError(value interface{}) []byte {
	j, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return j
}

func TestEndToEnd(t *testing.T) {
	db, _ := buntdb.Open(":memory:")
	a := NewApp(db)
	fillDb(db)

	tests := []struct {
		parallel   bool
		method     string
		path       string
		reqBody    io.Reader
		wantStatus int
		wantBody   []byte
	}{
		{method: "GET", path: "/teammates", reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(Teammates), parallel: true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.method, tt.path), func(t *testing.T) {
			if tt.parallel {
				t.Parallel()
			}

			req, _ := http.NewRequest(tt.method, tt.path, tt.reqBody)
			rr := httptest.NewRecorder()
			a.endpoints.Router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("%s %s -> status %d, expected %d", tt.method, tt.path, rr.Code, tt.wantStatus)
			}

			if body := rr.Body.Bytes(); !bytes.Equal(body, tt.wantBody) {
				t.Errorf("%s %s -> body \n %s \n expected \n %s", tt.method, tt.path, body, tt.wantBody)
			}
		})
	}
}
