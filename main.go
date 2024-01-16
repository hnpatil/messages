package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/external/db"
	"github.com/hnpatil/messages/handler"
	"github.com/hnpatil/messages/repository"
	"github.com/hnpatil/messages/usecase"
	"github.com/hnpatil/messages/utils/api"
	"github.com/hnpatil/messages/utils/config"
	"github.com/hnpatil/messages/utils/logger"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(config.GetConfig),

		fx.Provide(db.GetInstance),
		fx.Provide(logger.GetInstance),

		fx.Provide(api.NewEngine),
		fx.Provide(api.NewApiRoute),

		repositories(),
		usecases(),

		fx.Provide(handler.NewHandler),

		fx.Invoke(func(handler *handler.Handler, engine *gin.Engine) {
			handler.RegisterRoutes()

			err := engine.Run()
			if err != nil {
				panic(err)
			}
		}),
	).Run()
}

func usecases() fx.Option {
	return fx.Provide(
		usecase.NewAuth,
		usecase.NewMessages,
		usecase.NewUsers,
	)
}

func repositories() fx.Option {
	return fx.Provide(
		repository.NewAuthCodes,
	)
}
