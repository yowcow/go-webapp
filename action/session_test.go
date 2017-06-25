package action

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func build() *gin.Engine {
	router := gin.New()
	store := sessions.NewCookieStore([]byte("hogefuga"))
	router.Use(sessions.Sessions("mysession", store))
	return router
}

func TestHandleSetSession(t *testing.T) {
	assert := assert.New(t)

	query := url.Values{}
	query.Add("val", "hogehoge")
	body := bytes.NewBufferString(query.Encode())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	router := build()
	router.POST("/", HandleSetSession)
	router.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(200, resp.StatusCode)

	resbody := &struct {
		Success bool `json:"success"`
	}{}
	decoder := json.NewDecoder(resp.Body)

	assert.Nil(decoder.Decode(resbody))
	assert.True(resbody.Success)

	cookies := resp.Cookies()

	assert.Equal(1, len(cookies))

	assert.Equal("mysession", cookies[0].Name)
	assert.True(len(cookies[0].Value) > 0)
}

func TestHandleGetSession(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("cookie", "mysession=MTQ5ODEzMjQyMXxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQVVBQTNaaGJBWnpkSEpwYm1jTUNnQUlhRzluWldodloyVT18Zyaa5yY0NjGCd7jWkXWLT5wMJTFkaCetheXQOyQdxuU=;")

	router := build()
	router.GET("/", HandleGetSession)
	router.ServeHTTP(w, req)

	resp := w.Result()

	assert.Equal(200, resp.StatusCode)
	assert.Equal("hogehoge", resp.Body.String())
}
