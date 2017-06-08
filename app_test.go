package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestRoot(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("User-Agent", "MyBrowser/1")
	w := httptest.NewRecorder()

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Next()
		c.Header("X-Middleware", "hoge")
	})
	Build(r)
	r.ServeHTTP(w, req)

	respdata := map[string]interface{}{}
	json.Unmarshal(w.Body.Bytes(), &respdata)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("Gin", w.Header().Get("x-powered-by"))
	assert.Equal("hoge", w.Header().Get("x-middleware"))
	assert.Equal("world", respdata["hello"])
	assert.Equal("MyBrowser/1", respdata["ua"])
}

func TestJsonBody(t *testing.T) {
	assert := assert.New(t)

	jsonb, _ := json.Marshal(map[string]interface{}{
		"version": 1234,
	})
	reqbody := bytes.NewBuffer(jsonb)

	req := httptest.NewRequest("POST", "/json", reqbody)
	req.Header.Add("content-type", "application/json")
	w := httptest.NewRecorder()

	r := gin.New()
	Build(r)
	r.ServeHTTP(w, req)

	respdata := struct {
		Ver int `json:"version"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &respdata)

	assert.Equal(1234, respdata.Ver)
}

func TestFormBody(t *testing.T) {
	assert := assert.New(t)

	q := url.Values{}
	q.Set("hello", "=&world")
	q.Set("version", "2345")
	reqbody := bytes.NewBufferString(q.Encode())

	req := httptest.NewRequest("POST", "/form", reqbody)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	r := gin.New()
	Build(r)
	r.ServeHTTP(w, req)

	respdata := struct {
		Ver int `json:"version"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &respdata)

	assert.Equal(2345, respdata.Ver)
}

func TestLoginFails(t *testing.T) {
	assert := assert.New(t)

	jsonb, _ := json.Marshal(struct {
		Id     int    `json:"id"`
		Passwd string `json:"password"`
	}{111, "hogefuga"})
	reqbody := bytes.NewBuffer(jsonb)

	req := httptest.NewRequest("POST", "/login", reqbody)
	req.Header.Add("content-type", "application/json")
	w := httptest.NewRecorder()

	r := gin.New()
	Build(r)
	r.ServeHTTP(w, req)

	respdata := struct {
		Success bool `json:"success"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &respdata)

	assert.Equal(http.StatusForbidden, w.Code)
	assert.Equal(false, respdata.Success)
}

func TestLoginSucceeds(t *testing.T) {
	assert := assert.New(t)

	reqbody := bytes.NewBuffer([]byte(`
		{
			"id": 12345,
			"password": "mypassword"
		}
	`))

	req := httptest.NewRequest("POST", "/login", reqbody)
	req.Header.Add("content-type", "application/json")
	w := httptest.NewRecorder()

	r := gin.New()
	Build(r)
	r.ServeHTTP(w, req)

	respdata := struct {
		Success bool `json:"success"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &respdata)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(true, respdata.Success)
}

func TestSetSession(t *testing.T) {
	assert := assert.New(t)

	q := url.Values{}
	q.Set("val", "hogefuga")
	reqbody := bytes.NewBufferString(q.Encode())

	req := httptest.NewRequest("POST", "/session", reqbody)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	r := gin.New()
	Build(r)
	r.ServeHTTP(w, req)

	respdata := struct {
		Success bool `json:"success"`
	}{}
	json.Unmarshal(w.Body.Bytes(), &respdata)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(true, respdata.Success)
	assert.True(0 < len(w.Header().Get("set-cookie")))
}

func TestGetSession(t *testing.T) {
	assert := assert.New(t)

	req := httptest.NewRequest("GET", "/session", nil)
	req.Header.Add("cookie", "mysession=MTQ5NjkyMzY3OHxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQVVBQTNaaGJBWnpkSEpwYm1jTUNnQUlhRzluWldaMVoyRT18K3KFPKR05z2Ke6soM4rkr9KDHo0TwLUU9RQS7wRIJ0o=;")
	w := httptest.NewRecorder()

	r := gin.New()
	Build(r)
	r.ServeHTTP(w, req)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("hogefuga", w.Body.String())
}
