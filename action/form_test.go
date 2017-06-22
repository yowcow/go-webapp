package action

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleFormBody(t *testing.T) {
	assert := assert.New(t)

	query := url.Values{}
	query.Set("hello", "=&world")
	query.Set("version", "2345")
	body := bytes.NewBufferString(query.Encode())

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	router := gin.New()
	router.POST("/", HandleFormBody)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)

	res := &struct {
		V int `json:"version"`
	}{}
	decoder := json.NewDecoder(w.Body)

	assert.Nil(decoder.Decode(res))
	assert.Equal(2345, res.V)
}

func TestMultipartFormBody(t *testing.T) {
	assert := assert.New(t)

	body := &bytes.Buffer{}
	bufwriter := multipart.NewWriter(body)

	fr, _ := os.Open("./static/hello.html")
	fw, _ := bufwriter.CreateFormFile("myupload", "hello.html")
	io.Copy(fw, fr)

	fw, _ = bufwriter.CreateFormField("hello")
	fw.Write([]byte("こんにちは"))

	bufwriter.Close()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/", body)
	req.Header.Add("content-type", bufwriter.FormDataContentType())

	router := gin.New()
	router.POST("/", HandleMultipartFormBody)
	router.ServeHTTP(w, req)

	assert.Equal(200, w.Code)

	res := &struct {
		Filename string `json:"filename"`
		Hello    string `json:"hello"`
	}{}
	decoder := json.NewDecoder(w.Body)

	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Nil(decoder.Decode(res))
	assert.Equal("hello.html", res.Filename)
	assert.Equal("こんにちは", res.Hello)
}
