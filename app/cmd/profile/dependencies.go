package profile

import (
	service "github.com/tipee/account/app/services"
	"github.com/tipee/account/db"
	"github.com/tipee/account/db/repositories"
	"github.com/tipee/account/pkg/configs"
	"go.uber.org/dig"
)

func initDependecies() *dig.Container {
	configs.InitConfig()
	c := dig.New()

	//database
	_ = c.Provide(db.New)
	//repositories
	_ = c.Provide(repositories.New)
	//service
	_ = c.Provide(service.InitService)
	return c
}
