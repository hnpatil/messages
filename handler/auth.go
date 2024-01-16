package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/usecase"
	"github.com/loopfz/gadgeto/tonic"
)

func (handler *Handler) registerAuthRoutes() {
	auth := handler.api.Group("/auth")

	auth.POST("/code", tonic.Handler(handler.handleCreateAuthCode, http.StatusCreated))
	auth.POST("/token", tonic.Handler(handler.handleCreateAuthToken, http.StatusOK))
	auth.POST("/refresh", tonic.Handler(handler.handleRefreshAuthToken, http.StatusOK))
}

type CreateAuthCodeInput struct {
	Identifier string `json:"identifier" validate:"required"`
}

func (handler *Handler) handleCreateAuthCode(ctx *gin.Context, input *CreateAuthCodeInput) error {
	return handler.auth.GenerateAuthCode(usecase.NewContext(ctx), input.Identifier)
}

type CreateAuthTokenInput struct {
	AuthCode string `json:"auth_code" validate:"required"`
}

func (handler *Handler) handleCreateAuthToken(ctx *gin.Context, input *CreateAuthTokenInput) (*usecase.Token, error) {
	return handler.auth.GetAuthToken(usecase.NewContext(ctx), input.AuthCode)
}

type RefreshAuthTokenInput struct {
	RefrehToken string `json:"refreh_token" valudate:"required"`
}

func (handler *Handler) handleRefreshAuthToken(ctx *gin.Context, input *RefreshAuthTokenInput) (*usecase.Token, error) {
	return handler.auth.RefereshToken(usecase.NewContext(ctx), input.RefrehToken)
}
