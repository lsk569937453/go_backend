package task_controller_package

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbPingSuccess(t *testing.T) {
	router := gin.New()
	const path = "/userRoleList"
	router.GET("/api/db/dbPing", DbPing)
	_, statusCode, err := Get("/api/db/dbPing", router)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, statusCode)
}
