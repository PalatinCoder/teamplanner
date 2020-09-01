package endpoints

import (
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
