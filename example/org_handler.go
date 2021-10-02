package example

import (
	"net/http"

	"github.com/uptrace/bun-starter-kit/bunapp"
	"github.com/uptrace/bunrouter"
)

type OrgHandler struct {
	app *bunapp.App
}

func NewOrgHandler(app *bunapp.App) *OrgHandler {
	return &OrgHandler{
		app: app,
	}
}

func (h *OrgHandler) List(w http.ResponseWriter, req bunrouter.Request) error {
	ctx := req.Context()

	var orgs []Org
	if err := h.app.DB().NewSelect().Model(&orgs).Relation("Owner").Scan(ctx); err != nil {
		return err
	}

	return bunrouter.JSON(w, bunrouter.H{
		"orgs": orgs,
	})
}
