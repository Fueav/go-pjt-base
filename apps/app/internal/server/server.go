package server

import (
	"go-pjt-base/apps/app/config"
	"go-pjt-base/apps/app/internal/router"
	"go-pjt-base/pkg/commands"
	"go-pjt-base/pkg/common/xgin"
)

type server struct {
	ginServer *xgin.GinServer
	cfg       *config.Config
}

func NewServer() commands.MainInstance {
	return &server{cfg: config.NewConfig()}
}

func (s *server) Initialize() (err error) {
	s.ginServer = xgin.NewGinServer()
	router.Register(s.ginServer.Engine)
	return
}

func (s *server) RunLoop() {
	s.ginServer.Run(s.cfg.Port)
}

func (s *server) Destroy() {

}
