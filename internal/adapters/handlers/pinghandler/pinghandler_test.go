package pinghandler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/companies/internal/adapters/handlers/pinghandler"
	"github.com/stretchr/testify/assert"
)

func TestHandler_New(t *testing.T) {
	handler := pinghandler.New()

	assert.Equal(t, http.MethodGet, handler.GetHttpMethod())
	assert.Equal(t, "/ping", handler.GetRelativePath())
}

func TestHandler_ServeHTTP(t *testing.T) {
	handler := pinghandler.New()

	request := httptest.NewRequest("GET", "/ping", nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)

	body, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, "pong", string(body))
	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}
