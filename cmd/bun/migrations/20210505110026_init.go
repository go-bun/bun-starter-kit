package migrations

import (
	"context"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun-starter-kit/bunapp"
	"github.com/uptrace/bun-starter-kit/example"
	"github.com/uptrace/bun/fixture"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		app := bunapp.AppFromContext(ctx)

		db.RegisterModel((*example.User)(nil), (*example.Org)(nil))

		loader := fixture.NewLoader(db, fixture.WithRecreateTables())
		return loader.Load(ctx, os.DirFS(app.Path("example/testdata")), "fixture.yaml")
	}, nil)
}
