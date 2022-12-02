package randall

import (
	"io"
	"net/http"
)

type httpRequestProvider interface {
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}
