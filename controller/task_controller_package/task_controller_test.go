package task_controller_package

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_backend/util"
	"go_backend/vojo"
	"math/rand"
	"testing"
)

func TestTaskGet(t *testing.T) {
	router := gin.New()
	const path = "/api/task/getAll"
	router.GET(path, TaskGet)
	byteArr, resStatusCode, err := util.Get(path, router)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res.Rescode)
}
func TestTaskGetByUserIdSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/task/getByUserId"
	router.POST(path, TaskGet)
	reqBody := "{\"user_id\":\"-1\"}"
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res.Rescode)
}
func TestTaskGetByIdSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/task/getById"
	router.POST(path, TaskGetById)
	reqBody := "{\"id\":30}"
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, res.Rescode)
}
func TestTaskAddSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/task/add"
	router.POST(path, TaskAdd)
	reqBody := "{\"name\":\"1\",\"cron_expression\":\"*/5 * * * * ?\",\"url\":\"http://45.32.63.93:9001/api/test/getTest\"}"
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.NORMAL_RESPONSE_STATUS, res.Rescode)
}
func TestTaskAddFail(t *testing.T) {
	router := gin.New()
	const path = "/api/task/add"
	router.POST(path, TaskAdd)
	reqBody := "{\"name\":\"1\",\"cron_expression\":\"*/5 * * * * ?\",\"url\":\"http://45.32.63.93:9394/api/search\"}"
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.ERROR_STATUS_PARAM_WRONG, res.Rescode)
}
func TestTaskUpdateSuccess(t *testing.T) {
	router := gin.New()
	const path = "/api/task/updateById"
	router.POST(path, TaskUpdate)
	rand.Seed(2)
	number := rand.Intn(1000 * 1000)

	reqBody := "{\"id\":30,\"cron_expression\":\"lsk%d\",\"url\":\"test\"}"
	reqBody = fmt.Sprintf(reqBody, number)
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.NORMAL_RESPONSE_STATUS, res.Rescode)
}
func TestTaskUpdateFail(t *testing.T) {
	router := gin.New()
	const path = "/api/task/updateById"
	router.POST(path, TaskUpdate)
	reqBody := "{\"cron_expression\":\"lsk\",\"url\":\"aaatest\"}"
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.ERROR_RESPONSE_STATUS, res.Rescode)
}
func TestTaskDeleteFail(t *testing.T) {
	router := gin.New()
	const path = "/api/task/delById"
	router.POST(path, TaskDelete)
	reqBody := "{\"id\":100}"
	byteArr, resStatusCode, err := util.Post(path, router, reqBody)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resStatusCode)
	var res vojo.BaseRes
	err = json.Unmarshal(byteArr, &res)
	assert.Equal(t, nil, err)
	assert.Equal(t, vojo.ERROR_RESPONSE_STATUS, res.Rescode)
}
