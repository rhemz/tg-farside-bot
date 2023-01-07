package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	if assert.NoError(t, HealthEndpoint(c)) {
		assert.Equal(t, "{\n  \"message\": \"ok\"\n}\n", rec.Body.String())
	}
}
