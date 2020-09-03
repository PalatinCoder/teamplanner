package endpoints

import (
	"errors"
	"log"
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

func (e *Endpoints) setMatch(res http.ResponseWriter, req *http.Request) {
	var match model.Match
	err := decodeJSONBody(res, req, &match)
	if err != nil {
		var e ErrMalformedRequest
		if errors.Is(err, &e) {
			http.Error(res, e.msg, e.code)
		} else {
			log.Println(err.Error())
			respondWithError(res, http.StatusInternalServerError, err.Error())
		}
		return
	}

	replaced, err := e.repo.SetMatch(&match)
	if err != nil {
		switch {
		case err.Error() == "invalid":
			respondWithError(res, http.StatusBadRequest, "invalid match object")
		case err.Error() == "not found":
			respondWithError(res, http.StatusNotFound, err.Error())
		default:
			respondWithError(res, http.StatusInternalServerError, err.Error())
		}
		return
	}

	status := http.StatusCreated
	if replaced {
		status = http.StatusOK
	}

	respondWithJSON(res, status, &match)
}
