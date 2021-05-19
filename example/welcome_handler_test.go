package example_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/uptrace/bun-starter-kit/example"
	"github.com/uptrace/bun-starter-kit/testbed"
)

func TestWelcomeHandler(t *testing.T) {
	_, app := testbed.StartApp(t)
	defer app.Stop()

	handler := example.NewWelcomeHandler(app)

	{
		req := testbed.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()

		err := handler.Welcome(resp, req)
		require.NoError(t, err)
		require.Contains(t, resp.Body.String(), "Welcome")
	}

	{
		req := testbed.NewRequest("GET", "/hello", nil)
		resp := httptest.NewRecorder()

		err := handler.Hello(resp, req)
		require.NoError(t, err)
		require.Contains(t, resp.Body.String(), "Hello")
	}
}
