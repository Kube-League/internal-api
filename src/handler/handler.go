package handler

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"internal-api/src/utils"
	"io"
	"net/http"
)

type handler struct {
	Id string
	W  http.ResponseWriter
}

func Root(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "handler.Root", W: w}
	h.respond("Hey :3")
}

func (h *handler) getId() string {
	return "handler." + h.Id
}

func (h *handler) notNil(err error) {
	l := utils.Log{Id: h.getId()}
	l.Error(err)
	h.W.WriteHeader(http.StatusInternalServerError)
}

func (h *handler) badMethod() {
	h.respondCode(http.StatusMethodNotAllowed, "Method not allowed")
}

func (h *handler) respond(data any) {
	h.respondCode(http.StatusOK, data)
}

func (h *handler) respondCode(code int, data any) {
	res := Response{
		Code: uint(code),
		Data: data,
	}
	if code != http.StatusOK {
		h.W.WriteHeader(code)
	}
	b, err := res.Json()
	if err != nil {
		h.notNil(err)
		return
	}
	h.W.Write(b)
}

func (h *handler) query(db *gorm.DB, dest interface{}, condition string, args ...string) error {
	err := db.First(dest).Where(condition, args).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return nil
	}
	return err
}

func (h *handler) found(err error) bool {
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return false
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		h.respondCode(http.StatusNotFound, h.Id+" not found")
		return false
	}
	return true
}

func generateAsk(dest interface{}, r *http.Request) bool {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}
	err = json.Unmarshal(b, dest)
	if err != nil {
		l := utils.Log{Id: "handler.generateAsk"}
		l.Error(err)
		return false
	}
	return true
}
