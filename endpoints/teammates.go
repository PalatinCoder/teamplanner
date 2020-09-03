package endpoints

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"jan-sl.de/teamplanner/model"
)

func (e *Endpoints) getTeammates(res http.ResponseWriter, req *http.Request) {
	mates, err := e.repo.GetTeammates()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, mates)
}

func (e *Endpoints) getTeammate(res http.ResponseWriter, req *http.Request) {
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
	respondWithJSON(res, http.StatusOK, mate)
}

func (e *Endpoints) setTeammate(res http.ResponseWriter, req *http.Request) {
	var mate model.Teammate
	err := decodeJSONBody(res, req, &mate)
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

	replaced, err := e.repo.SetTeammate(&mate)
	if err != nil {
		switch {
		case err.Error() == "invalid":
			respondWithError(res, http.StatusBadRequest, "invalid teammate object")
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

	respondWithJSON(res, status, &mate)
}
