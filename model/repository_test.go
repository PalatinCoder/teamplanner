package model

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tidwall/buntdb"
)

var emptyDb, filledDb *buntdb.DB

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

func TestMain(m *testing.M) {

	/* Setup test dbs */
	emptyDb, _ = buntdb.Open(":memory:")
	createIndexes(emptyDb)
	filledDb, _ = buntdb.Open(":memory:")
	createIndexes(filledDb)
	if err := filledDb.Update(func(tx *buntdb.Tx) error {
		for _, m := range Teammates {
			j, _ := json.Marshal(m)
			tx.Set(m.Key(), string(j), nil)
		}
		for _, m := range Matches {
			j, _ := json.Marshal(m)
			tx.Set(m.Key(), string(j), nil)
		}
		for _, m := range Votes {
			j, _ := json.Marshal(m)
			tx.Set(m.Key(), string(j), nil)
		}

		return nil
	}); err != nil {
		panic(err)
	}
	m.Run()
}

func TestNewRepository(t *testing.T) {
	type args struct {
		db *buntdb.DB
	}
	tests := []struct {
		name string
		args args
		want *BuntDb
	}{
		{name: "creates a repository with the given db", args: args{emptyDb}, want: &BuntDb{emptyDb}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBuntDb(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("creates the indexes", func(t *testing.T) {
		want := []string{"matches", "teammates", "votes"}
		if got, _ := NewBuntDb(emptyDb).db.Indexes(); !cmp.Equal(got, want) {
			t.Errorf("NewRepository() -> db.Indexes() = %v, want %v", got, want)
		}
	})
}

func TestRepository_GetTeammates(t *testing.T) {
	type fields struct {
		db *buntdb.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Teammate
		wantErr bool
	}{
		{name: "empty db gives nil", fields: fields{db: emptyDb}, want: nil, wantErr: false}, // nil slice != empty slice!
		{name: "filled db gives all items", fields: fields{db: filledDb}, want: Teammates, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BuntDb{
				db: tt.fields.db,
			}
			got, err := r.GetTeammates()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetTeammates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Repository.GetTeammates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetMatches(t *testing.T) {
	type fields struct {
		db *buntdb.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Match
		wantErr bool
	}{
		{name: "empty db gives nil", fields: fields{db: emptyDb}, want: nil, wantErr: false}, // nil slice != empty slice!
		{name: "filled db gives all items", fields: fields{db: filledDb}, want: Matches, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BuntDb{
				db: tt.fields.db,
			}
			got, err := r.GetMatches()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetMatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Repository.GetMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetVotes(t *testing.T) {
	type fields struct {
		db *buntdb.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Vote
		wantErr bool
	}{
		{name: "empty db gives nil", fields: fields{db: emptyDb}, want: nil, wantErr: false}, // nil slice != empty slice!
		{name: "filled db gives all items", fields: fields{db: filledDb}, want: Votes, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BuntDb{
				db: tt.fields.db,
			}
			got, err := r.GetVotes()
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetVotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Repository.GetVotes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetVotesByTeammate(t *testing.T) {
	type fields struct {
		db *buntdb.DB
	}
	type args struct {
		mate Teammate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Vote
		wantErr bool
	}{
		{name: "empty db gives nil", fields: fields{db: emptyDb}, args: args{mate: Teammates[0]}, want: nil, wantErr: false}, // nil slice != empty slice!
		//TODO {name: "empty teammate gives error", fields: fields{db: emptyDb}, args: args{mate: Teammate{}}, want: nil, wantErr: true},
		{
			name:    "filled db gives all items for teammate",
			fields:  fields{db: filledDb},
			args:    args{mate: Teammates[0]},
			want:    filterVotes(Votes, func(v Vote) bool { return cmp.Equal(v.Teammate, Teammates[0]) }),
			wantErr: false,
		},
		{
			name:    "filled db gives all items for another teammate",
			fields:  fields{db: filledDb},
			args:    args{mate: Teammates[1]},
			want:    filterVotes(Votes, func(v Vote) bool { return cmp.Equal(v.Teammate, Teammates[1]) }),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BuntDb{
				db: tt.fields.db,
			}
			got, err := r.GetVotesByTeammate(tt.args.mate)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetVotesByTeammate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Repository.GetVotesByTeammate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetVotesForMatch(t *testing.T) {
	type fields struct {
		db *buntdb.DB
	}
	type args struct {
		match Match
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Vote
		wantErr bool
	}{
		{name: "empty db gives nil", fields: fields{db: emptyDb}, args: args{match: Matches[2]}, want: nil, wantErr: false}, // nil slice != empty slice!
		//TODO {name: "empty teammate gives error", fields: fields{db: emptyDb}, args: args{mate: Teammate{}}, want: nil, wantErr: true},
		{
			name:    "filled db gives all items for teammate",
			fields:  fields{db: filledDb},
			args:    args{match: Matches[2]},
			want:    filterVotes(Votes, func(v Vote) bool { return cmp.Equal(v.Match, Matches[2]) }),
			wantErr: false,
		},
		{
			name:    "filled db gives all items for another teammate",
			fields:  fields{db: filledDb},
			args:    args{match: Matches[3]},
			want:    filterVotes(Votes, func(v Vote) bool { return cmp.Equal(v.Match, Matches[3]) }),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &BuntDb{
				db: tt.fields.db,
			}
			got, err := r.GetVotesForMatch(tt.args.match)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetVotesForMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Repository.GetVotesForMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
