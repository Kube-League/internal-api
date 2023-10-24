package routes

import (
	"encoding/json"
	"internal-api/src/db/sql"
	"internal-api/src/utils"
	"io"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	respond("routes.Root", "Hey :3", w)
}

func Team(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		badMethod(w)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		notNil("routes.Team", err, w)
		return
	}
	var info sql.TeamInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		notNil("routes.Team", err, w)
		return
	}
	id := sql.CreateTeam(info)
	respond("routes.Team", id, w)
}

func Teams(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		badMethod(w)
		return
	}
	var teams []sql.Team
	sql.DB.Find(&teams)
	respond("routes.Teams", teams, w)
}

func notNil(id string, err error, w http.ResponseWriter) {
	utils.LogError(id, err)
	w.WriteHeader(http.StatusInternalServerError)
}

func badMethod(w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	res := Response{
		Code: 200,
		Data: "Method not allowed",
	}
	b, err := res.Json()
	if err != nil {
		notNil("routes.Team", err, w)
		return
	}
	w.Write(b)
}

func respond(id string, data any, w http.ResponseWriter) {
	res := Response{
		Code: 200,
		Data: data,
	}
	b, err := res.Json()
	if err != nil {
		notNil(id, err, w)
		return
	}
	w.Write(b)
}
