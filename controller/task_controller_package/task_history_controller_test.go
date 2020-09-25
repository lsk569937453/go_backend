package task_controller_package

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_backend/vojo"
	"testing"
)

func TestTaskHistoryGetByTaskIdFail(t *testing.T) {
	router := gin.New()
	const path = "/api/taskHistory/getByTaskId"
	router.POST(path, TaskHistoryGetByTaskId)
	reqBody := "{\"task_id\":30}"
	byteArr, resStatusCode, err := Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.ERROR_RESPONSE_STATUS, res.Rescode)
}

func TestTaskHistoryGetByPageFail(t *testing.T) {
	router := gin.New()
	const path = "/api/taskHistory/getByPage"
	router.POST(path, TaskHistoryGetByPage)
	reqBody := "{\"task_id\":41,\"page\":{\"id\":58948,\"pageSize\":20}}"
	byteArr, resStatusCode, err := Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.ERROR_RESPONSE_STATUS, res.Rescode)
}
