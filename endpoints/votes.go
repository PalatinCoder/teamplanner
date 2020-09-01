package endpoints

import "net/http"

func (e *Endpoints) getVotes(res http.ResponseWriter, req *http.Request) {
	votes, err := e.repo.GetVotes()
	if err != nil {
		respondWithError(res, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(res, http.StatusOK, votes)
}

func (e *Endpoints) getVotesForTeammate(res http.ResponseWriter, req *http.Request) {

}
