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

type askPlayer struct {
	Results bool `json:"result"`
}

func Player(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Player", W: w}
	if r.Method != http.MethodPost {
		h.badMethod()
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.notNil(err)
		return
	}
	var info sql.PlayerInfo
	err = json.Unmarshal(b, &info)
	if err != nil {
		h.notNil(err)
		return
	}
	id, err := sql.CreatePlayer(info)
	if err != nil {
		//TODO: handle same data
		h.notNil(err)
		return
	}
	h.respondCode(http.StatusCreated, id)
}

func Players(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Players", W: w}
	allowed := []string{http.MethodGet, http.MethodDelete}
	if !utils.ArrStringContains(allowed, r.Method) {
		h.badMethod()
		return
	}
	switch r.Method {
	case http.MethodGet:
		var players []sql.Player
		var err error
		var a askPlayer
		ok := generateAsk(&a, r)
		if ok {
			err = a.preload().Find(&players).Error
		} else {
			err = sql.DB.Find(&players).Error
		}
		if !h.found(err) {
			return
		}
		h.respond(players)
	case http.MethodDelete:
		err := sql.DB.Exec("DELETE FROM players").Error
		if err != nil {
			h.notNil(err)
			return
		}
		h.respond("Deleted")
	}
}

func PlayerId(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "PlayerId", W: w}
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
	var player sql.Player
	var err error
	var a askPlayer
	ok = generateAsk(&a, r)
	if ok {
		err = a.preload().First(&player, id).Error
	} else {
		err = sql.DB.First(&player, id).Error
	}
	if !h.found(err) {
		return
	}
	h.respond(player)
}

func PlayerName(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "PlayerName", W: w}
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
	var player sql.Player
	var err error
	var a askPlayer
	ok = generateAsk(&a, r)
	if ok {
		err = h.query(a.preload(), &player, "last_name = ?", name)
	} else {
		err = h.query(sql.DB, &player, "last_name = ?", name)
	}
	if !h.found(err) {
		return
	}
	h.respond(player)
}

func PlayerUuid(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "PlayerName", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	p := mux.Vars(r)
	uuid, ok := p["uuid"]
	if !ok {
		h.respondCode(http.StatusBadRequest, "Missing uuid")
		return
	}
	var player sql.Player
	var err error
	var a askPlayer
	ok = generateAsk(&a, r)
	if ok {
		err = h.query(a.preload(), &player, "uuid = ?", uuid)
	} else {
		err = h.query(sql.DB, &player, "uuid = ?", uuid)
	}
	if !h.found(err) {
		return
	}
	h.respond(player)
}

func PlayerDiscord(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "PlayerName", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	p := mux.Vars(r)
	id, ok := p["id"]
	if !ok {
		h.respondCode(http.StatusBadRequest, "Missing id")
		return
	}
	var player sql.Player
	var err error
	var a askPlayer
	ok = generateAsk(&a, r)
	if ok {
		err = h.query(a.preload(), &player, "discord_id = ?", id)
	} else {
		err = h.query(sql.DB, &player, "discord_id = ?", id)
	}
	if !h.found(err) {
		return
	}
	h.respond(player)
}

func (a *askPlayer) preload() *gorm.DB {
	b := sql.DB.Model(&sql.Player{})
	if a.Results {
		b = b.Preload("Results")
	}
	return b
}
