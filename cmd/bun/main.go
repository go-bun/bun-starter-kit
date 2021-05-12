package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/uptrace/bun-starter-kit/app"
	"github.com/uptrace/bun-starter-kit/cmd/bun/migrations"
	_ "github.com/uptrace/bun-starter-kit/example"
	"github.com/uptrace/bun-starter-kit/httputil"
	"github.com/uptrace/bun/migrate"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "bun",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "environment",
			},
		},
		Commands: []*cli.Command{
			serverCommand,
			dbCommand,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

var serverCommand = &cli.Command{
	Name:  "runserver",
	Usage: "start API server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "addr",
			Value: ":8000",
			Usage: "serve address",
		},
	},
	Action: func(c *cli.Context) error {
		myapp, err := app.Start(c.Context, "api", c.String("env"))
		if err != nil {
			return err
		}
		defer myapp.Stop()

		var handler http.Handler
		handler = myapp.Router()
		handler = httputil.ExitOnPanicHandler{Next: handler}

		srv := &http.Server{
			Addr:         c.String("addr"),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
			Handler:      handler,
		}
		go func() {
			if err := srv.ListenAndServe(); err != nil && !isServerClosed(err) {
				log.Printf("ListenAndServe failed: %s", err)
			}
		}()

		fmt.Printf("listening on %s\n", srv.Addr)
		fmt.Println(app.WaitExitSignal())

		return srv.Shutdown(c.Context)
	},
}

var dbCommand = &cli.Command{
	Name:  "db",
	Usage: "manage database migrations",
	Subcommands: []*cli.Command{
		{
			Name:  "init",
			Usage: "create migration tables",
			Action: func(c *cli.Context) error {
				app, migrator := migrator(c)
				defer app.Stop()

				return migrator.Init(c.Context, app.DB())
			},
		},
		{
			Name:  "migrate",
			Usage: "migrate database",
			Action: func(c *cli.Context) error {
				app, migrator := migrator(c)
				defer app.Stop()

				return migrator.Migrate(c.Context, app.DB())
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback the last migration group",
			Action: func(c *cli.Context) error {
				app, migrator := migrator(c)
				defer app.Stop()

				return migrator.Rollback(c.Context, app.DB())
			},
		},
		{
			Name:  "unlock",
			Usage: "unlock migrations",
			Action: func(c *cli.Context) error {
				app, migrator := migrator(c)
				defer app.Stop()

				return migrator.Unlock(c.Context, app.DB())
			},
		},
		{
			Name:  "create_go",
			Usage: "create a Go migration",
			Action: func(c *cli.Context) error {
				app, migrator := migrator(c)
				defer app.Stop()

				return migrator.CreateGo(c.Context, app.DB(), c.Args().Get(0))
			},
		},
		{
			Name:  "create_sql",
			Usage: "create a SQL migration",
			Action: func(c *cli.Context) error {
				app, migrator := migrator(c)
				defer app.Stop()

				return migrator.CreateSQL(c.Context, app.DB(), c.Args().Get(0))
			},
		},
	},
}

func migrator(c *cli.Context) (*app.App, *migrate.Migrator) {
	app, err := app.Start(c.Context, "api", c.String("env"))
	if err != nil {
		log.Fatal(err)
	}
	return app, migrations.Migrator
}

func isServerClosed(err error) bool {
	return err.Error() == "http: Server closed"
}
