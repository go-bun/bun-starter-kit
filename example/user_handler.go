package example

import (
	"net/http"

	"github.com/uptrace/bun-starter-kit/bunapp"
	"github.com/vmihailenco/treemux"
)

type UserHandler struct {
	app *bunapp.App
}

func NewUserHandler(app *bunapp.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) List(w http.ResponseWriter, req treemux.Request) error {
	ctx := req.Context()

	var users []User
	if err := h.app.DB().NewSelect().Model(&users).Scan(ctx); err != nil {
		return err
	}

	return treemux.JSON(w, treemux.H{
		"users": users,
	})
}
