package util

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"strings"
)

// Get send url
func Get(uri string, router *gin.Engine) ([]byte, int, error) {
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)

	return body, result.StatusCode, err
}

// Get send url
func Post(uri string, router *gin.Engine, reqBody string) ([]byte, int, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(strings.NewReader(reqBody))
	req := httptest.NewRequest("POST", uri, buf)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)

	return body, result.StatusCode, err
}
