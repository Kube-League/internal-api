package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"internal-api/src/db/sql"
	"internal-api/src/utils"
	"io"
	"net/http"
)

type askTeam struct {
	Matches bool `json:"matches"`
	Players bool `json:"players"`
}

func Team(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Team", W: w}
	if r.Method != http.MethodPost {
		h.badMethod()
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.notNil(err)
		return
	}
	var info sql.TeamInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		h.notNil(err)
		return
	}
	id, err := sql.CreateTeam(info)
	if err != nil {
		//TODO: handle same name
		h.notNil(err)
		return
	}
	h.respondCode(http.StatusCreated, id)
}

func Teams(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Teams", W: w}
	allowed := []string{http.MethodGet, http.MethodDelete}
	if !utils.ArrStringContains(allowed, r.Method) {
		h.badMethod()
		return
	}
	switch r.Method {
	case http.MethodGet:
		var teams []sql.Team
		var err error
		var a askTeam
		ok := generateAsk(&a, r)
		if ok {
			err = a.preload().Find(&teams).Error
		} else {
			err = sql.DB.Find(&teams).Error
		}
		if !h.found(err) {
			return
		}
		h.respond(teams)
	case http.MethodDelete:
		err := sql.DB.Exec("DELETE FROM teams").Error
		if err != nil {
			h.notNil(err)
			return
		}
		h.respond("Deleted")
	}
}

func TeamId(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "TeamId", W: w}
	allowed := []string{http.MethodGet, http.MethodPut, http.MethodDelete}
	if !utils.ArrStringContains(allowed, r.Method) {
		h.badMethod()
		return
	}
	p := mux.Vars(r)
	id, ok := p["id"]
	if !ok {
		h.respondCode(http.StatusBadRequest, "Missing id")
		return
	}
	var team sql.Team
	var err error
	var a askTeam
	ok = generateAsk(&a, r)
	if ok {
		err = a.preload().First(&team, id).Error
	} else {
		err = sql.DB.First(&team, id).Error
	}
	if !h.found(err) {
		return
	}
	h.respond(team)
}

func TeamName(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "TeamName", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	p := mux.Vars(r)
	name, ok := p["name"]
	if !ok {
		h.respondCode(http.StatusBadRequest, "Missing name")
		return
	}
	var team sql.Team
	var err error
	var a askTeam
	ok = generateAsk(&a, r)
	if ok {
		err = h.query(a.preload(), &team, "name = ?", name)
	} else {
		err = h.query(sql.DB, &team, "name = ?", name)
	}
	if !h.found(err) {
		return
	}
	h.respond(team)
}

func (a *askTeam) preload() *gorm.DB {
	b := sql.DB.Model(&sql.Team{})
	if a.Matches {
		b = b.Preload("Matches").Preload("MatchesWon").Preload("Results")
	}
	if a.Players {
		b = b.Preload("Players")
	}
	return b
}
