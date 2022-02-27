package example

import (
	"context"

	"github.com/go-bun/bun-starter-kit/bunapp"
)

func init() {
	bunapp.OnStart("example.init", func(ctx context.Context, app *bunapp.App) error {
		router := app.Router()
		api := app.APIRouter()

		welcomeHandler := NewWelcomeHandler(app)
		userHandler := NewUserHandler(app)
		orgHandler := NewOrgHandler(app)

		router.GET("/", welcomeHandler.Welcome)
		router.GET("/hello", welcomeHandler.Hello)

		api.GET("/users", userHandler.List)
		api.GET("/orgs", orgHandler.List)

		return nil
	})
}
