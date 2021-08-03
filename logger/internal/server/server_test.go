package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	ts := httptest.NewServer(SetupRouter())
	defer ts.Close()

	// set up request
	newreq := func(method, url string, body io.Reader) *http.Request {
		r, err := http.NewRequest(method, url, body)
		if err != nil {
			t.Fatal(err)
		}
		return r
	}

	// for now just testing status
	// will need to refactor this section when we add more handlers
	tests := []struct {
		name string
		r    *http.Request
	}{
		{name: "testing status check", r: newreq("GET", ts.URL+"/status", nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Do(tt.r)
			defer resp.Body.Close()

			if err != nil {
				t.Fatal(err)
			}

			require.Equal(t, resp.StatusCode, http.StatusOK)
		})
	}
}
