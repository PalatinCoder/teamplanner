package endpoints

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"jan-sl.de/teamplanner/model"
)

func (e *Endpoints) getMatches(res http.ResponseWriter, req *http.Request) {
	matches, err := e.repo.GetMatches()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, matches)
}

func (e *Endpoints) getMatch(res http.ResponseWriter, req *http.Request) {
	date, err := time.Parse("20060102", mux.Vars(req)["id"])
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
	}
	match := model.Match{Date: date}
	err = e.repo.GetMatch(&match)
	if err != nil {
		switch err.Error() {
		case "not found": // This isn't really nice...
			respondWithError(res, http.StatusNotFound, err.Error())
		default:
			respondWithError(res, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(res, http.StatusOK, match)
}
