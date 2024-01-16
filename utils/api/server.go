package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/utils/config"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func NewEngine(logger logrus.FieldLogger) *gin.Engine {
	tonic.SetErrorHook(tonicErrorHook)
	engine := gin.New()

	engine.Use(loggerMiddleware(logger))
	engine.Use(ginlogrus.Logger(logger), gin.Recovery())

	return engine
}

func NewApiRoute(cfg *config.Config, engine *gin.Engine) *gin.RouterGroup {
	return engine.Group(fmt.Sprintf("/api/%s", cfg.GetValue(config.VERSION)))
}

func loggerMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := logger.WithFields(logrus.Fields{
			"path":      ctx.Request.URL.Path,
			"method":    ctx.Request.Method,
			"hostname":  ctx.Request.Host,
			"client_ip": ctx.ClientIP(),
		})

		ctx.Set("logger", lg)
		ctx.Next()
	}
}
