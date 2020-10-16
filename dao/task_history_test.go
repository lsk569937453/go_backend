package dao

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go_backend/vojo"
	"os"
	"testing"
)

func TestHistoryInsertFail(t *testing.T) {
	history := &vojo.TasksHistory{}
	err := HistoryInsert(history)
	assert.Error(t, err)
}
func TestHistoryInsertSuccess(t *testing.T) {
	history := &vojo.TasksHistory{
		Task_id:     15,
		Exec_result: "req_result",
		Exec_code:   0,
		Exec_time:   "aaa",
	}
	err := HistoryInsert(history)
	assert.Equal(t, nil, err)
}
func TestMain(m *testing.M) {
	fmt.Println("begin")
	m.Run()
	fmt.Println("end")
	os.Exit(0)
}
