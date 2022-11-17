package example

import (
	"net/http"

	"github.com/go-bun/bun-starter-kit/bunapp"
	"github.com/uptrace/bunrouter"
)

type UserHandler struct {
	app *bunapp.App
}

func NewUserHandler(app *bunapp.App) *UserHandler {
	return &UserHandler{
		app: app,
	}
}

func (h *UserHandler) List(w http.ResponseWriter, req bunrouter.Request) error {
	ctx := req.Context()

	var users []User

	if err := h.app.DB().NewSelect().Model(&users).Scan(ctx); err != nil {
		return err
	}

	return bunrouter.JSON(w, bunrouter.H{
		"users": users,
	})
}

func (h *UserHandler) Get(w http.ResponseWriter, req bunrouter.Request) error {
	ctx := req.Context()

	id := req.Param("id")

	var user User
	if err := h.app.DB().NewSelect().Where("id = ?", id).Model(&user).Scan(ctx); err != nil {
		return err
	}

	return bunrouter.JSON(w, user)
}
