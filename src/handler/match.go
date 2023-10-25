package handler

import (
	"encoding/json"
	"internal-api/src/db/sql"
	"io"
	"net/http"
)

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
		h.notNil(err)
		return
	}
	h.respond(id)
}
