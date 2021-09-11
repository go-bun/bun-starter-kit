package example

import (
	"embed"
	"net/http"
	"text/template"

	"github.com/uptrace/bun-starter-kit/bunapp"
	"github.com/uptrace/treemux"
)

//go:embed templates/*
var templates embed.FS

type WelcomeHandler struct {
	app *bunapp.App
	tpl *template.Template
}

func NewWelcomeHandler(app *bunapp.App) *WelcomeHandler {
	tpl, err := template.New("").ParseFS(templates, "templates/*.html")
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

func (h *WelcomeHandler) Hello(w http.ResponseWriter, req treemux.Request) error {
	if err := h.tpl.ExecuteTemplate(w, "hello.html", nil); err != nil {
		return err
	}
	return nil
}
