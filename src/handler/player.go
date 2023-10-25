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
	h.respond(id)
}

func Players(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "Players", W: w}
	if r.Method != http.MethodGet {
		h.badMethod()
		return
	}
	var players []sql.Player
	err := sql.DB.Find(&players).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "No players found")
		return
	}
	h.respond(players)
}

func PlayerId(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "PlayerId", W: w}
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
	var player sql.Player
	err := sql.DB.First(&player, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Player not found")
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
	err := h.query(sql.DB, &player, "last_name = ?", name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Player not found")
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
	err := h.query(sql.DB, &player, "uuid = ?", uuid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Player not found")
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
	err := h.query(sql.DB, &player, "discord_id = ?", id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, "Player not found")
		return
	}
	h.respond(player)
}
