package file_share

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go_backend/util"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGetClientIDSuccess(t *testing.T) {
	router := gin.New()
	router.GET("/api/shareFile/getClientID", GetClientID)

	_, statusCode, err := util.Get("/api/shareFile/getClientID", router)
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, statusCode)
}
func TestIoCopySuccess(t *testing.T) {

	out, err := os.OpenFile("../../resource/test.file", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	assert.Equal(t, nil, err)
	wt := bufio.NewWriter(out)
	src := strings.NewReader("Nidhi: F\nRahul: M\nNisha: F\n")

	n, err := io.Copy(wt, src)
	fmt.Printf("write byte count:%d", n)
	wt.Flush()
	assert.Equal(t, nil, err)
}
