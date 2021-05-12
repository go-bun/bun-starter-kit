package app

import (
	"net/http"

	"github.com/uptrace/bun-starter-kit/httputil/httperror"
	"github.com/vmihailenco/treemux"
	"github.com/vmihailenco/treemux/extra/reqlog"
	"github.com/vmihailenco/treemux/extra/treemuxgzip"
	"github.com/vmihailenco/treemux/extra/treemuxotel"
)

func (app *App) initRouter() {
	opts := []treemux.Option{
		treemux.WithMiddleware(treemuxgzip.NewMiddleware()),
		treemux.WithMiddleware(treemuxotel.NewMiddleware()),
	}
	if app.IsDebug() {
		opts = append(opts, treemux.WithMiddleware(reqlog.NewMiddleware()))
	}
	opts = append(opts, treemux.WithMiddleware(app.errorHandler))

	app.router = treemux.New(opts...)
	app.apiRouter = app.router.NewGroup("/api",
		treemux.WithMiddleware(corsMiddleware),
	)
}

func (app *App) errorHandler(next treemux.HandlerFunc) treemux.HandlerFunc {
	return func(w http.ResponseWriter, req treemux.Request) error {
		err := next(w, req)
		if err == nil {
			return nil
		}

		httpErr := httperror.From(err, app.IsDebug())
		if httpErr.Status != 0 {
			w.WriteHeader(httpErr.Status)
		}
		_ = treemux.JSON(w, httpErr)

		return err
	}
}

func corsMiddleware(next treemux.HandlerFunc) treemux.HandlerFunc {
	return func(w http.ResponseWriter, req treemux.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			return next(w, req)
		}

		h := w.Header()

		h.Set("Access-Control-Allow-Origin", origin)
		h.Set("Access-Control-Allow-Credentials", "true")

		// CORS preflight.
		if req.Method == http.MethodOptions {
			h.Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,HEAD")
			h.Set("Access-Control-Allow-Headers", "authorization,content-type")
			h.Set("Access-Control-Max-Age", "86400")
			return nil
		}

		return next(w, req)
	}
}
