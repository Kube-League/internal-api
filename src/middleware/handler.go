package middleware

import (
	"errors"
	"gorm.io/gorm"
	"internal-api/src/db/sql"
	"internal-api/src/utils"
	"net/http"
)

type handler struct {
	Id string
	W  http.ResponseWriter
}

func Root(w http.ResponseWriter, r *http.Request) {
	h := handler{Id: "middleware.Root", W: w}
	h.respond("Hey :3")
}

func (h *handler) getId() string {
	return "middleware." + h.Id
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

func (h *handler) query(dest interface{}, condition string, args ...string) error {
	err := sql.DB.First(dest).Where(condition, args).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		h.notNil(err)
		return nil
	}
	return err
}