package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"jan-sl.de/teamplanner/model"
)

// ErrMalformedRequest describes an error related to a bad client request (body)
type ErrMalformedRequest struct {
	code int
	msg  string
}

func (e *ErrMalformedRequest) Error() string {
	return e.msg
}

// Endpoints holds the HTTP endpoints for the API
type Endpoints struct {
	repo   model.Dataprovider
	Router *mux.Router
}

// NewEndpoints creates a new set of HTTP endpoints on the given repository
func NewEndpoints(repo model.Dataprovider, router *mux.Router) *Endpoints {
	e := &Endpoints{repo, router}
	e.Router.HandleFunc("/teammates", e.getTeammates).Methods("GET")
	e.Router.HandleFunc("/teammate", e.setTeammate).Methods("POST")
	e.Router.HandleFunc("/teammate/{id:[0-9]+}", e.getTeammate).Methods("GET")
	e.Router.HandleFunc("/teammate/{id:[0-9]+}/votes", e.getVotesForTeammate).Methods("GET")
	e.Router.HandleFunc("/matches", e.getMatches).Methods("GET")
	e.Router.HandleFunc("/match", e.setMatch).Methods("POST")
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

func decodeJSONBody(res http.ResponseWriter, req *http.Request, dst model.Model) error {
	if contentType := req.Header.Get("Content-Type"); contentType != "application/json" {
		return &ErrMalformedRequest{http.StatusUnsupportedMediaType, "Content-Type is not application/json"}
	}
	req.Body = http.MaxBytesReader(res, req.Body, 1048576) // 1MB
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains baldy-formed JSON at %d", syntaxError.Offset)
			return &ErrMalformedRequest{code: http.StatusBadRequest, msg: msg}
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains baldy-formed JSON")
			return &ErrMalformedRequest{code: http.StatusBadRequest, msg: msg}
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains invalid value for field %q at %d", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &ErrMalformedRequest{code: http.StatusBadRequest, msg: msg}
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &ErrMalformedRequest{code: http.StatusBadRequest, msg: msg}
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &ErrMalformedRequest{http.StatusBadRequest, msg}
		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &ErrMalformedRequest{http.StatusRequestEntityTooLarge, msg}
		default:
			return err
		}
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return &ErrMalformedRequest{http.StatusBadRequest, "Request body must only contain a single object"}
	}
	return nil
}
