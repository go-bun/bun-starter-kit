package example

import (
	"context"
	"html/template"
	"net/http"

	"github.com/uptrace/bun-starter-kit/app"
	"github.com/vmihailenco/treemux"
)

func init() {
	app.OnStart("example.init", func(ctx context.Context, app *app.App) error {
		welcomeHandler := NewWelcomeHandler(app)

		app.Router().GET("/", welcomeHandler.Welcome)

		return nil
	})
}

type WelcomeHandler struct {
	app *app.App
	tpl *template.Template
}

func NewWelcomeHandler(app *app.App) *WelcomeHandler {
	tpl, err := template.New("").ParseGlob(app.Path("example", "templates", "*.html"))
	if err != nil {
		panic(err)
	}

	return &WelcomeHandler{
		app: app,
		tpl: tpl,
	}
}

func (h *WelcomeHandler) Welcome(w http.ResponseWriter, req treemux.Request) error {
	if err := h.tpl.ExecuteTemplate(w, "welcome.html", nil); err != nil {
		return err
	}
	return nil
}
