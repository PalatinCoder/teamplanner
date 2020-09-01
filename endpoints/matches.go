package endpoints

import "net/http"

func (e *Endpoints) getMatches(res http.ResponseWriter, req *http.Request) {
	matches, err := e.repo.GetMatches()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, matches)
}