package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/tidwall/buntdb"
	"jan-sl.de/teamplanner/endpoints"
	"jan-sl.de/teamplanner/model"
)

// App is the wrapper for everything
type App struct {
	repo      *model.BuntDb
	endpoints *endpoints.Endpoints
}

// NewApp initializes a new instance of the application
func NewApp(db *buntdb.DB) *App {
	app := &App{}
	app.repo = model.NewBuntDb(db)
	app.endpoints = endpoints.NewEndpoints(app.repo, mux.NewRouter())
	return app
}

// Run runs the HTTP server on the given address
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, handlers.CombinedLoggingHandler(os.Stdout, a.endpoints.Router)))
}

func main() {
	var dbfile, addr string
	var ok bool

	if dbfile, ok = os.LookupEnv("DBPATH"); !ok {
		dbfile = "/data/planner.db"
	}
	if addr, ok = os.LookupEnv("LISTENADDR"); !ok {
		addr = ":8042"
	}

	db, err := buntdb.Open(dbfile)
	if err != nil {
		panic(err)
	}

	app := NewApp(db)
	app.Run(addr)
}
