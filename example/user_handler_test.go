package example_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun-starter-kit/bunapp"
	"github.com/uptrace/bun-starter-kit/example"
	"github.com/uptrace/bun-starter-kit/testbed"
	"github.com/uptrace/bun/dbfixture"
)

func TestUserHandler(t *testing.T) {
	_, app := testbed.StartApp(t)
	defer app.Stop()

	fixture := loadFixture(t, app)
	testUser := fixture.MustRow("User.test").(*example.User)

	handler := example.NewUserHandler(app)

	{
		req := testbed.NewRequest("GET", "/api/users", nil)
		resp := httptest.NewRecorder()

		err := handler.List(resp, req)
		require.NoError(t, err)
		require.Contains(t, resp.Body.String(), testUser.Name)
	}
}

func loadFixture(t *testing.T, app *bunapp.App) *dbfixture.Fixture {
	db := app.DB()
	db.RegisterModel((*example.User)(nil), (*example.Org)(nil))

	fixture := dbfixture.New(db, dbfixture.WithRecreateTables())
	err := fixture.Load(app.Context(), bunapp.FS(), "fixture/fixture.yaml")
	require.NoError(t, err)

	return fixture
}
