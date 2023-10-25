package main

import (
	"github.com/gorilla/mux"
	"internal-api/src/config"
	"internal-api/src/db/sql"
	"internal-api/src/middleware"
	"log"
	"net/http"
	"time"
)

func main() {
	config.Setup()
	sql.Setup()

	err := sql.DB.AutoMigrate(&sql.Match{}) // it will also set up team, player, etc.
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", middleware.Root)
	// team
	r.HandleFunc("/team", middleware.Team)
	r.HandleFunc("/teams", middleware.Teams)
	r.HandleFunc("/team/{id:[0-9]+}", middleware.TeamId).
		Name("team (id)")
	r.HandleFunc("/team/{name}", middleware.TeamName).
		Name("team (name)")
	// player

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
