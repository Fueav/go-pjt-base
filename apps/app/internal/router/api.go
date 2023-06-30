package router

import (
	"github.com/gin-gonic/gin"
	logger "github.com/ipfs/go-log"
)

var log = logger.Logger("router")

func Register(engine *gin.Engine) {
	v1PublicGroup := engine.Group("api/v1")
	registerV1PublicRoutes(v1PublicGroup)

}

func registerV1PublicRoutes(group *gin.RouterGroup) {
	registerV1InscriptionsRouter(group)
}

func registerV1InscriptionsRouter(group *gin.RouterGroup) {
}
