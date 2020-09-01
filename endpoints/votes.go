package endpoints

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"jan-sl.de/teamplanner/model"
)

func (e *Endpoints) getVotes(res http.ResponseWriter, req *http.Request) {
	votes, err := e.repo.GetVotes()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, votes)
}

func (e *Endpoints) getVotesForTeammate(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
	}
	mate := model.Teammate{Position: id}
	err = e.repo.GetTeammate(&mate)
	if err != nil {
		switch err.Error() {
		case "not found":
			respondWithError(res, http.StatusNotFound, err.Error())
		default:
			respondWithError(res, http.StatusInternalServerError, err.Error())
		}
		return
	}
	votes, err := e.repo.GetVotesByTeammate(mate)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, votes)
}

func (e *Endpoints) getVotesForMatch(res http.ResponseWriter, req *http.Request) {
	date, err := time.Parse("20060102", mux.Vars(req)["id"])
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
	}
	match := model.Match{Date: date}
	err = e.repo.GetMatch(&match)
	if err != nil {
		switch err.Error() {
		case "not found":
			respondWithError(res, http.StatusNotFound, err.Error())
		default:
			respondWithError(res, http.StatusInternalServerError, err.Error())
		}
		return
	}
	votes, err := e.repo.GetVotesForMatch(match)
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, votes)
}
