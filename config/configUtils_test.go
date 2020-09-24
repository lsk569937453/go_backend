package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetValueSuccess(t *testing.T) {
	str, err := GetValue("redis", "address")

	assert.Equal(t, nil, err)

	assert.Equal(t, "127.0.0.1:6379", str)
}
func TestGetValueFail(t *testing.T) {
	_, err := GetValue("redis", "addressOther")

	assert.NotEqual(t, "127.0.0.1:6379", err)
}
