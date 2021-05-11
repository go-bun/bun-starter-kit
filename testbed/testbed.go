package testbed

import (
	"io"
	"net/http/httptest"

	"github.com/vmihailenco/treemux"
)

func NewRequest(method, target string, body io.Reader) treemux.Request {
	return treemux.NewRequest(httptest.NewRequest(method, target, body))
}
