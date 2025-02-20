package handler

import (
	"strings"

	"github.com/KeyZenDB/KeyZenDB/internal"
)

type Handler struct {
	command map[string]func(args []any) any
}

func NewHandler() *Handler {
	db := internal.NewDb()

	return &Handler{
		command: map[string]func(args []any) any{
			"PING":  func(_ []any) any { return db.Ping() },
			"SET":   func(data []any) any { return db.Set(string(data[1].(string)), data[2]) },
			"GET":   func(data []any) any { return db.Get(string(data[1].(string))) },
			"DEL":   func(data []any) any { return db.Del(string(data[1].(string))) },
			"CLEAR": func(_ []any) any { return db.Clear() },
		},
	}
}

func (h *Handler) ProcessCommand(cmd []any) string {
	command := strings.ToUpper(cmd[0].(string))
	h.command[command](cmd)
	return ""
}
