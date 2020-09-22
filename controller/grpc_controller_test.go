package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_backend/vojo"
	"testing"
)

func TestGrpcGetServiceListSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/grpc/getServiceList"
	router.POST(path, GrpcGetServiceList)
	reqBody := "{\"url\":\"127.0.0.1:9000\"}"
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
