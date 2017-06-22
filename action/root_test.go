package action

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleRoot(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("user-agent", "foobar")

	router := gin.New()
	router.GET("/", HandleRoot)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)
	assert.Equal("Gin", w.Header().Get("x-powered-by"))

	res := &struct {
		Hello string `json:"hello"`
		Ua    string `json:"ua"`
	}{}
	decoder := json.NewDecoder(w.Body)

	assert.Nil(decoder.Decode(res))
	assert.Equal("world", res.Hello)
	assert.Equal("foobar", res.Ua)
}
