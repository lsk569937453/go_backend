package dao

import (
	"github.com/stretchr/testify/assert"
	"go_backend/vojo"
	"testing"
)

func TestAddTaskFail(t *testing.T) {
	history := &vojo.TaskInsertReq{}
	err := AddTask(history)
	assert.Equal(t, int64(-1), err)
}
