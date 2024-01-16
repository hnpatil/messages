package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/entity"
	"github.com/hnpatil/messages/usecase"
	"github.com/hnpatil/messages/utils/api"
	"github.com/loopfz/gadgeto/tonic"
)

func (handler *Handler) registerMessagesRoutes() {
	messages := handler.api.Group("/messages")
	messages.Use(api.JwtAuthMiddleware(handler.auth))

	messages.POST("", tonic.Handler(handler.createMessage, http.StatusCreated))

	conversations := handler.api.Group("/conversations")
	conversations.Use(api.JwtAuthMiddleware(handler.auth))

	conversations.GET("", tonic.Handler(handler.listConversations, http.StatusOK))
	conversations.GET(":conv_id/messages", tonic.Handler(handler.listMessages, http.StatusOK))
}

type CreateMessageInput struct {
	Recipient string `json:"recipient" validate:"required"`
	Text      string `json:"text" validate:"required"`
}

func (h *Handler) createMessage(ctx *gin.Context, input *CreateMessageInput) (*entity.Message, error) {
	return h.messages.SendMessage(usecase.NewContext(ctx), input.Recipient, input.Text)
}

func (h *Handler) listConversations(ctx *gin.Context) ([]*entity.Conversation, error) {
	return h.messages.ListConversations(usecase.NewContext(ctx))
}

type ListMessagesInput struct {
	ConvID string `path:"conv_id" validate:"required"`
}

func (h *Handler) listMessages(ctx *gin.Context, input *ListMessagesInput) ([]*entity.Message, error) {
	return h.messages.ListMessages(usecase.NewContext(ctx), input.ConvID)
}
