package test

import (
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
	"icando/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck_BestCase(t *testing.T) {
	TestRunner(
		t, func(server *server.Server, db *gorm.DB, _ Fixture) {
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
