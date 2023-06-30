package main

import (
	logger "github.com/ipfs/go-log"
	"go-pjt-base/apps/app/config"
	"go-pjt-base/pkg/common/xmysql"
	"go-pjt-base/pkg/common/xredis"
)

var log = logger.Logger("main")

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	logger.SetAllLoggers(logger.LevelInfo)

}
