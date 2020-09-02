package model

import (
	"time"
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
}
