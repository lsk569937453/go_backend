package dao

import (
	"github.com/stretchr/testify/assert"
	"go_backend/vojo"
	"testing"
)

func TestAddTaskSuccess(t *testing.T) {
	taskInsertReq := &vojo.TaskInsertReq{}
	err := AddTask(taskInsertReq)
	assert.NotEqual(t, int64(-1), err)
}
