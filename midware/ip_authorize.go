package midware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IpAuthorize() gin.HandlerFunc{
	return func(c *gin.Context){
		clientIp:=c.ClientIP()

		if clientIp!="192.168.0.1" {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		}
	}
}



