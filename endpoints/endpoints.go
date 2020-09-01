package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"jan-sl.de/teamplanner/model"
)

// Endpoints holds the HTTP endpoints for the API
type Endpoints struct {
	repo   model.Dataprovider
	Router *mux.Router
}

// NewEndpoints creates a new set of HTTP endpoints on the given repository
func NewEndpoints(repo model.Dataprovider, router *mux.Router) *Endpoints {
	e := &Endpoints{repo, router}
	e.Router.HandleFunc("/teammates", e.getTeammates).Methods("GET")
	e.Router.HandleFunc("/teammate/{id:[0-9]+}", e.getTeammate).Methods("GET")
	e.Router.HandleFunc("/teammate/{id:[0-9]+}/votes", e.getVotesForTeammate).Methods("GET")
	e.Router.HandleFunc("/matches", e.getMatches).Methods("GET")
	e.Router.HandleFunc("/match/{id:[0-9]{8}}/votes", e.getVotesForMatch).Methods("GET")
	e.Router.HandleFunc("/match/{id:[0-9]{8}}", e.getMatch).Methods("GET")
	e.Router.HandleFunc("/votes", e.getVotes).Methods("GET")
	return e
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(res)
}
