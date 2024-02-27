package test

import (
	"github.com/stretchr/testify/require"
	"icando/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck_BestCase(t *testing.T) {
	testRunner(
		t, func(server *server.Server) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				"GET",
				"/",
				nil,
			)
			server.Engine.ServeHTTP(w, req)
			require.Equal(t, 200, w.Code)
		},
	)
}
