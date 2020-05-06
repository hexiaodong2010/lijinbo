package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin()  gin.HandlerFunc{
	return func(c *gin.Context) {
		if openid:=c.GetHeader("openid"); len(openid) == 0 {
			c.String(http.StatusForbidden, "please login")
			c.Abort()
			return
		}
	}
}
