package main

import (
	logger "github.com/ipfs/go-log"
	"go-pjt-base/apps/app/config"
	"go-pjt-base/apps/app/internal/server"
	"go-pjt-base/pkg/commands"
	"go-pjt-base/pkg/common/xmysql"
	"go-pjt-base/pkg/common/xredis"
)

func init() {
	conf := config.GetConfig()
	xmysql.NewMysqlClient(conf.Mysql)
	xredis.NewRedisClient(conf.Redis)
}

func main() {
	logger.SetLogLevel("*", "INFO")
	commands.Run(server.NewServer())
}
