package main

import (
	"github.com/gorilla/mux"
	"internal-api/src/config"
	"internal-api/src/db/sql"
	"internal-api/src/handler"
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

	r.HandleFunc("/", handler.Root)
	// team
	r.HandleFunc("/team", handler.Team)
	r.HandleFunc("/teams", handler.Teams)
	r.HandleFunc("/team/id/{id:[0-9]+}", handler.TeamId).
		Name("team (id)")
	r.HandleFunc("/team/name/{name}", handler.TeamName).
		Name("team (name)")
	// player
	r.HandleFunc("/player", handler.Player)
	r.HandleFunc("/players", handler.Players)
	r.HandleFunc("/player/id/{id:[0-9]+}", handler.PlayerId).
		Name("player (id)")
	r.HandleFunc("/player/uuid/{uuid:\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b}", handler.PlayerUuid).
		Name("player (uuid)")
	r.HandleFunc("/player/discord/{id:[0-9]{18}", handler.PlayerDiscord).
		Name("player (discord)")
	r.HandleFunc("/player/name/{name}", handler.PlayerName).
		Name("player (name)")
	// match
	r.HandleFunc("/match", handler.Match)
	r.HandleFunc("/matches", handler.Matches)
	r.HandleFunc("/match/id/{id:[0-9]+}", handler.MatchId).
		Name("match (id)")
	r.HandleFunc("/match/time/{time:[0-9]+}", handler.MatchTime).
		Name("match (name)")
	// match result
	r.HandleFunc("/match/id/{id:[0-9]+}/results", handler.MatchResult).
		Name("match results (id)")
	r.HandleFunc("/match/results", handler.MatchResultCreate).
		Name("match results create")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
