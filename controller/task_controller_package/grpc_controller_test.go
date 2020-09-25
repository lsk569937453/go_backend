package task_controller_package

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_backend/vojo"
	"testing"
)

func TestGrpcGetServiceListSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/grpc/getServiceList"
	router.POST(path, GrpcGetServiceList)
	reqBody := "{\"url\":\"45.32.63.93:9000\"}"
	_, statusCode, err := Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, statusCode)
}
func TestGrpcGetServiceListFailed(t *testing.T) {
	router := gin.New()
	const path = "/api/grpc/getServiceList"
	router.POST(path, GrpcGetServiceList)
	reqBody := "{\"url\":\"45.32.63.93:9001\"}"
	byteArr, resStatusCode, err := Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, -1, res.Rescode)
}
func TestGrpcRemoteInvokeSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/grpc/remoteInvoke"
	router.POST(path, GrpcRemoteInvoke)
	reqBody := "{\"url\":\"45.32.63.93:9000\",\"serviceName\":\"test.MaxSize\",\"methodName\":\"Echo\",\"reqJson\":\"{}\"}"
	byteArr, resStatusCode, err := Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res.Rescode)
}
func TestGrpcRemoteInvokeFail(t *testing.T) {
	router := gin.New()
	const path = "/api/grpc/remoteInvoke"
	router.POST(path, GrpcRemoteInvoke)
	reqBody := "{\"url\":\"45.32.63.93:9001\",\"serviceName\":\"test.MaxSize\",\"methodName\":\"Echo\",\"reqJson\":\"{}\"}"
	byteArr, resStatusCode, err := Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, -1, res.Rescode)
}
func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
}
