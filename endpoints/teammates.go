package endpoints

import "net/http"

func (e *Endpoints) getTeammates(res http.ResponseWriter, req *http.Request) {
	mates, err := e.repo.GetTeammates()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, mates)
}
