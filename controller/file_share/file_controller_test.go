package file_share

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_backend/util"
	"testing"
)

func TestGetClientIDSuccess(t *testing.T) {
	router := gin.New()
	router.GET("/api/shareFile/getClientID", GetClientID)

	_, statusCode, err := util.Get("/api/shareFile/getClientID", router)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, statusCode)
}
