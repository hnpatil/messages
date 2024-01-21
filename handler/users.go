package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/handler/response"
	"github.com/hnpatil/messages/usecase"
	"github.com/hnpatil/messages/utils/api"
	"github.com/loopfz/gadgeto/tonic"
)

func (handler *Handler) registerUsersRoutes() {
	users := handler.api.Group("/users")

	users.Use(api.JwtAuthMiddleware(handler.auth))
	users.POST("", tonic.Handler(handler.handleCreateUser, http.StatusCreated))
}

type CreateUserInput struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (handler *Handler) handleCreateUser(c *gin.Context, input *CreateUserInput) (*response.User, error) {
	res, err := handler.users.CreateUser(usecase.NewContext(c), input.FirstName, input.LastName)
	if err != nil {
		return nil, err
	}

	return response.ToUserResponse(res), nil
}
