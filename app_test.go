package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStatic(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/static/hello.html", nil)

	router := gin.New()
	Build(router)
	router.ServeHTTP(w, req)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("<h1>Hello</h1>\n", w.Body.String())
}
