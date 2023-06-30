package dig

import (
	"go-pjt-base/apps/app/config"
	"go.uber.org/dig"
)

var container = dig.New()

func init() {
	container.Provide(config.NewConfig)
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
