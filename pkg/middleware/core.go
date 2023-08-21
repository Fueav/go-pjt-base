package middleware

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
	"net/http"
	"runtime/debug"
)

var log = logger.Logger("middleware")

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,Content-Type")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("Panic info is: %s", string(debug.Stack()))
			}
		}()

		c.Next()
	}
}
