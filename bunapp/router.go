package bunapp

import (
	"net/http"

	"github.com/uptrace/bun-starter-kit/httputil/httperror"
	"github.com/uptrace/bunrouter"
	"github.com/uptrace/bunrouter/extra/bunrouterotel"
	"github.com/uptrace/bunrouter/extra/reqlog"
)

func (app *App) initRouter() {
	app.router = bunrouter.New(
		bunrouter.WithMiddleware(bunrouterotel.NewMiddleware()),
		bunrouter.WithMiddleware(reqlog.NewMiddleware(
			reqlog.WithEnabled(app.IsDebug()),
			reqlog.FromEnv(""),
		)),
	)

	app.apiRouter = app.router.NewGroup("/api",
		bunrouter.WithMiddleware(corsMiddleware),
		bunrouter.WithMiddleware(app.errorHandler),
	)
}

func (app *App) errorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)
		if err == nil {
			return nil
		}

		httpErr := httperror.From(err, app.IsDebug())
		if httpErr.Status != 0 {
			w.WriteHeader(httpErr.Status)
		}
		_ = bunrouter.JSON(w, httpErr)

		return err
	}
}

func corsMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
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
