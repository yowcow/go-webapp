package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func request(method, url string, body io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	req := httptest.NewRequest(method, url, body)
	w := httptest.NewRecorder()
	return w, req
}

func build(middlewares ...func(*gin.Context)) *gin.Engine {
	routing := gin.New()
	for _, f := range middlewares {
		routing.Use(f)
	}
	Build(routing)
	return routing
}

func TestRoot(t *testing.T) {
	assert := assert.New(t)

	w, req := request("GET", "/", nil)
	req.Header.Add("user-Agent", "MyBrowser/1")

	r := build(func(c *gin.Context) {
		c.Header("X-Middleware", "hoge")
		c.Next()
	})
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

	w, req := request("POST", "/json", reqbody)
	req.Header.Add("content-type", "application/json")

	r := build()
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

	w, req := request("POST", "/form", reqbody)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	r := build()
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

	w, req := request("POST", "/login", reqbody)
	req.Header.Add("content-type", "application/json")

	r := build()
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

	w, req := request("POST", "/login", reqbody)
	req.Header.Add("content-type", "application/json")

	r := build()
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

	w, req := request("POST", "/session", reqbody)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	r := build()
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

	w, req := request("GET", "/session", nil)
	req.Header.Add("cookie", "mysession=MTQ5NjkyMzY3OHxEdi1CQkFFQ180SUFBUkFCRUFBQUpmLUNBQUVHYzNSeWFXNW5EQVVBQTNaaGJBWnpkSEpwYm1jTUNnQUlhRzluWldaMVoyRT18K3KFPKR05z2Ke6soM4rkr9KDHo0TwLUU9RQS7wRIJ0o=;")

	r := build()
	r.ServeHTTP(w, req)

	assert.Equal(http.StatusOK, w.Code)
	assert.Equal("hogefuga", w.Body.String())
}
