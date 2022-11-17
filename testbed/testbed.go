package testbed

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-bun/bun-starter-kit/bunapp"
	"github.com/stretchr/testify/require"
)

func NewRequest(method, target string, body io.Reader) *http.Request {
	return httptest.NewRequest(method, target, body)
}

func StartApp(t *testing.T) (context.Context, *bunapp.App) {
	ctx, app, err := bunapp.Start(context.TODO(), "test", "test")
	require.NoError(t, err)
	return ctx, app
}
