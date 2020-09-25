package util

import (
	"fmt"
	"testing"
	"time"
)

func TestGetCurrentTime(t *testing.T) {
	for i := 0; i <= 10; i++ {
		ti := GetCurrentTime()

		ss := fmt.Sprintf("ssss:%s", ti)
		fmt.Println(ss)
		time.Sleep(time.Second)

	}

	//assert.NotEqual(t, nil, ti)
}
