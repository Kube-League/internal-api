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

type askMatch struct {
	Results      bool `json:"result"`
	Participants bool `json:"participants"`
}

func Match(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Match", W: w}
	if r.Method != http.MethodPost {
		h.badMethod()
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.notNil(err)
		return
	}
	var info sql.MatchInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		h.notNil(err)
		return
	}
	id, err := sql.CreateMatch(info)
	if err != nil {
		//TODO: handle cannot find team
		h.notNil(err)
		return
	}
	h.respond(id)
}

func Matches(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Matches", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	var matches []sql.Match
	var err error
	var a askMatch
	ok := generateAsk(&a, r)
	if ok {
		err = a.preload().Find(&matches).Error
	} else {
		err = sql.DB.Find(&matches).Error
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "No players found")
		return
	}
	h.respond(matches)
}

func MatchId(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "MatchId", W: w}
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
	var match sql.Match
	var err error
	var a askMatch
	ok = generateAsk(&a, r)
	if ok {
		err = a.preload().First(&match, id).Error
	} else {
		err = sql.DB.First(&match, id).Error
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Match not found")
		return
	}
	h.respond(match)
}

func MatchTime(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "MatchTime", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	p := mux.Vars(r)
	t, ok := p["time"]
	if !ok {
		h.respondCode(http.StatusBadRequest, "Missing time")
		return
	}
	var match sql.Match
	var err error
	var a askMatch
	ok = generateAsk(&a, r)
	if ok {
		err = h.query(a.preload(), &match, "date = ?", t)
	} else {
		err = h.query(sql.DB, &match, "date = ?", t)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Player not found")
		return
	}
	h.respond(match)
}

func (a *askMatch) preload() *gorm.DB {
	b := sql.DB.Model(&sql.Team{})
	if a.Results {
		b = b.Preload("Results").Preload("PlayerMatchResult")
	}
	if a.Participants {
		b = b.Preload("Participants")
	}
	return b
}
