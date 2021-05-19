package example_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun-starter-kit/example"
	"github.com/uptrace/bun-starter-kit/testbed"
)

func TestOrgHandler(t *testing.T) {
	_, app := testbed.StartApp(t)
	defer app.Stop()

	fixture := loadFixture(t, app)
	testUser := fixture.MustRow("User.test").(*example.User)
	myOrg := fixture.MustRow("Org.my").(*example.Org)

	handler := example.NewOrgHandler(app)

	{
		req := testbed.NewRequest("GET", "/api/orgs", nil)
		resp := httptest.NewRecorder()

		err := handler.List(resp, req)
		require.NoError(t, err)
		require.Contains(t, resp.Body.String(), testUser.Name)
		require.Contains(t, resp.Body.String(), myOrg.Name)
	}
}
