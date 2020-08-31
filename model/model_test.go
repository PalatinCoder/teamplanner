package model

import (
	"testing"
	"time"
)

func TestMatch_Key(t *testing.T) {
	type fields struct {
		Date        time.Time
		Description string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "correct key for match", fields: fields{Date: time.Date(2020, 8, 30, 16, 31, 25, 60, time.UTC), Description: "Testdate"}, want: "match:20200830"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Match{
				Date:        tt.fields.Date,
				Description: tt.fields.Description,
			}
			if got := m.Key(); got != tt.want {
				t.Errorf("Match.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTeammate_Key(t *testing.T) {
	type fields struct {
		Name     string
		Position int
		Status   TeammateStatus
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "correct key for teammate", fields: fields{Name: "Tester", Position: 1, Status: StatusAvail}, want: "mate:1"},
		{name: "correct key for teammate", fields: fields{Name: "Tester", Position: 2, Status: StatusAvail}, want: "mate:2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Teammate{
				Name:     tt.fields.Name,
				Position: tt.fields.Position,
				Status:   tt.fields.Status,
			}
			if got := m.Key(); got != tt.want {
				t.Errorf("Teammate.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVote_Key(t *testing.T) {
	type fields struct {
		Teammate Teammate
		Match    Match
		Vote     VoteOption
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "correct key for vote", fields: fields{Teammate: Teammate{Position: 1}, Match: Match{Date: time.Date(2020, 8, 30, 16, 31, 25, 60, time.UTC)}, Vote: VoteYes}, want: "vote:1:20200830"},
		{name: "correct key for vote", fields: fields{Teammate: Teammate{Position: 1}, Match: Match{Date: time.Date(2020, 8, 30, 16, 31, 25, 60, time.UTC)}, Vote: VoteNo}, want: "vote:1:20200830"},
		{name: "correct key for vote", fields: fields{Teammate: Teammate{Position: 1}, Match: Match{Date: time.Date(2020, 8, 30, 16, 31, 25, 60, time.UTC)}, Vote: VoteMaybe}, want: "vote:1:20200830"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Vote{
				Teammate: tt.fields.Teammate,
				Match:    tt.fields.Match,
				Vote:     tt.fields.Vote,
			}
			if got := m.Key(); got != tt.want {
				t.Errorf("Vote.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
