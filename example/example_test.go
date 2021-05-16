package example_test

import (
	"context"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun-starter-kit/bunapp"
	"github.com/uptrace/bun-starter-kit/example"
	"github.com/uptrace/bun-starter-kit/testbed"
	"github.com/uptrace/bun/fixture"
)

var ctx = context.Background()

func TestUsers(t *testing.T) {
	app := startTestApp(t)
	defer app.Stop()

	var users []example.User
	err := app.DB().NewSelect().
		Model(&users).
		OrderExpr("id ASC").
		Scan(ctx)
	require.NoError(t, err)
	require.Len(t, users, 2)
	require.Equal(t, "test user", users[0].Name)
	require.Equal(t, "another user", users[1].Name)
}

func TestOrgs(t *testing.T) {
	app := startTestApp(t)
	defer app.Stop()

	var orgs []example.Org
	err := app.DB().NewSelect().
		Model(&orgs).
		Relation("Owner").
		OrderExpr("org.id ASC").
		Scan(ctx)
	require.NoError(t, err)
	require.Len(t, orgs, 2)

	org := orgs[0]
	require.Equal(t, "My Org", org.Name)
	require.Equal(t, "test user", org.Owner.Name)
}

func TestOrg(t *testing.T) {
	app := startTestApp(t)
	defer app.Stop()

	myorg := app.fixture.MustGet("Org", "my_org").(*example.Org)

	org := new(example.Org)
	err := app.DB().NewSelect().
		Model(org).
		Where("id = ?", myorg.ID).
		Scan(ctx)
	require.NoError(t, err)
	require.Equal(t, *org, *myorg)
}

func TestHandler(t *testing.T) {
	app := startTestApp(t)
	defer app.Stop()

	handler := example.NewWelcomeHandler(app.App)

	req := testbed.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	err := handler.Welcome(resp, req)
	require.NoError(t, err)
	require.Contains(t, resp.Body.String(), "Welcome")
}

type App struct {
	*bunapp.App
	fixture *fixture.Loader
}

func startTestApp(t *testing.T) App {
	ctx, app, err := bunapp.Start(context.TODO(), "test", "test")
	require.NoError(t, err)

	db := app.DB()
	db.RegisterModel((*example.User)(nil), (*example.Org)(nil))

	loader := fixture.NewLoader(db, fixture.WithDropTables())
	err = loader.Load(ctx, os.DirFS("testdata"), "fixture.yaml")
	require.NoError(t, err)

	return App{
		App:     app,
		fixture: loader,
	}
}
