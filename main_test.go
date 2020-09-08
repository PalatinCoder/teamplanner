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
		for _, m := range Matches {
			if err := insert(tx, m); err != nil {
				return err
			}
		}
		for _, m := range Votes {
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

func TestEndToEnd(t *testing.T) {
	db, _ := buntdb.Open(":memory:")
	a := NewApp(db)
	fillDb(db)

	tests := []struct {
		method     string
		path       string
		reqBody    io.Reader
		wantStatus int
		wantBody   []byte
	}{
		{method: "GET", path: "/teammates", reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(Teammates)},
		{method: "GET", path: fmt.Sprintf("/teammate/%s", Teammates[0].ID()), reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(Teammates[0])},
		{method: "GET", path: fmt.Sprintf("/teammate/%s/votes", Teammates[0].ID()), reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(filterVotes(Votes, func(v Vote) bool { return v.Teammate.ID() == Teammates[0].ID() }))},
		{method: "GET", path: "/matches", reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(Matches)},
		{method: "GET", path: fmt.Sprintf("/match/%s", Matches[1].ID()), reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(Matches[1])},
		{method: "GET", path: fmt.Sprintf("/match/%s/votes", Matches[1].ID()), reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(filterVotes(Votes, func(v Vote) bool { return v.Match.ID() == Matches[1].ID() }))},
		{method: "GET", path: "/votes", reqBody: nil, wantStatus: 200, wantBody: marshallJSONWithoutError(Votes)},
		{method: "POST", path: "/teammate", reqBody: bytes.NewReader(marshallJSONWithoutError(Teammates[1])), wantStatus: 200, wantBody: marshallJSONWithoutError(Teammates[1])},
		{method: "POST", path: "/match", reqBody: bytes.NewReader(marshallJSONWithoutError(Matches[1])), wantStatus: 200, wantBody: marshallJSONWithoutError(Matches[1])},
		{method: "POST", path: "/vote", reqBody: bytes.NewReader(marshallJSONWithoutError(Votes[1])), wantStatus: 200, wantBody: marshallJSONWithoutError(Votes[1])},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.method, tt.path), func(t *testing.T) {
			req, _ := http.NewRequest(tt.method, tt.path, tt.reqBody)
			rr := httptest.NewRecorder()
			if req.ContentLength > 0 {
				req.Header.Add("content-type", "application/json")
			}
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
