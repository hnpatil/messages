package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/usecase"
)

type Handler struct {
	api      *gin.RouterGroup
	auth     usecase.Auth
	users    usecase.Users
	messages usecase.Messages
}

func NewHandler(
	api *gin.RouterGroup,
	auth usecase.Auth,
	users usecase.Users,
	messages usecase.Messages,
) *Handler {
	return &Handler{
		api:      api,
		auth:     auth,
		users:    users,
		messages: messages,
	}
}

func (h *Handler) RegisterRoutes() {
	h.registerAuthRoutes()
	h.registerMessagesRoutes()
	h.registerUsersRoutes()
}
