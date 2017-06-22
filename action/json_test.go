package action

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleJsonBody(t *testing.T) {
	assert := assert.New(t)

	body := &bytes.Buffer{}
	encoder := json.NewEncoder(body)
	encoder.Encode(struct {
		V int `json:"version"`
	}{1234})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)

	router := gin.New()
	router.POST("/", HandleJsonBody)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)

	res := &struct {
		V int `json:"version"`
	}{}
	decoder := json.NewDecoder(w.Body)

	assert.Nil(decoder.Decode(res))
	assert.Equal(1234, res.V)
}
