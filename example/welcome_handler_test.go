package example_test

import (
	"net/http/httptest"
	"testing"

	"github.com/go-bun/bun-starter-kit/testbed"
	"github.com/stretchr/testify/require"
)

func TestWelcomeHandler(t *testing.T) {
	_, app := testbed.StartApp(t)
	defer app.Stop()

	{
		req := testbed.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()

		err := app.Router().ServeHTTPError(resp, req)

		require.NoError(t, err)
		require.Contains(t, resp.Body.String(), "Welcome")
	}

	{
		req := testbed.NewRequest("GET", "/hello", nil)
		resp := httptest.NewRecorder()

		err := app.Router().ServeHTTPError(resp, req)

		require.NoError(t, err)
		require.Contains(t, resp.Body.String(), "Hello")
	}
}
