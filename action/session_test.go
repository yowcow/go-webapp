package action

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"regexp"
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

	assert.Equal(200, w.Code)

	res := &struct {
		Success bool `json:"success"`
	}{}
	decoder := json.NewDecoder(w.Body)

	assert.Nil(decoder.Decode(res))
	assert.True(res.Success)

	re := regexp.MustCompile("^mysession=(?:[a-zA-Z0-9\\-\\_]+)=;.+")
	cookie := w.Header().Get("set-cookie")

	assert.True(re.MatchString(cookie))
}

func TestHandleGetSession(t *testing.T) {
	assert := assert.New(t)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("cookie", "mysession=MTQ5ODEzMjQyMXxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQVVBQTNaaGJBWnpkSEpwYm1jTUNnQUlhRzluWldodloyVT18Zyaa5yY0NjGCd7jWkXWLT5wMJTFkaCetheXQOyQdxuU=;")

	router := build()
	router.GET("/", HandleGetSession)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)
	assert.Equal("hogehoge", w.Body.String())
}
