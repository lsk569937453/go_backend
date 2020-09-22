package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http/httptest"
	"strings"
)

// Get send url
func Get(uri string, router *gin.Engine) ([]byte, int, error) {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, err := ioutil.ReadAll(result.Body)

	return body, result.StatusCode, err
}

// Get send url
func Post(uri string, router *gin.Engine, reqBody string) ([]byte, int, error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(strings.NewReader(reqBody))
	// 构造get请求
	req := httptest.NewRequest("POST", uri, buf)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body, err := ioutil.ReadAll(result.Body)

	return body, result.StatusCode, err
}
