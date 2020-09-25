package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSetSuccess(t *testing.T) {
	err := Set("c", "bssss")
	assert.Equal(t, nil, err)
}
func TestSetNxSuccess(t *testing.T) {
	err := SetNX("cxxx", "bssss", time.Hour)
	assert.Equal(t, nil, err)
}
