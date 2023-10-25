package handler

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"internal-api/src/db/sql"
	"internal-api/src/utils"
	"io"
	"net/http"
)

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
	h.respond(id)
}

func Teams(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Teams", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	var teams []sql.Team
	sql.DB.Find(&teams)
	h.respond(teams)
}

func TeamId(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "TeamId", W: w}
	allowed := []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch}
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
	err := sql.DB.First(&team, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Team not found")
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
	err := sql.DB.First(&team).Where("name = ?", name).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Team not found")
		return
	}
	h.respond(team)
}
