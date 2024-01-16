package repository

import (
	"context"
	"time"

	"github.com/hnpatil/messages/entity"
)

type AuthCodes interface {
	CreateAuthCode(ctx context.Context, identifier string, code string, expiresAt *time.Time) (*entity.AuthCode, error)
	GetAuthCode(ctx context.Context, code string) (*entity.AuthCode, error)
}

type Users interface {
	CreateUser(ctx context.Context, firstName string, lastName string, email string) (*entity.User, error)
	GetUser(ctx context.Context, email string) (*entity.User, error)
}

type Messages interface {
	CreateMessage(ctx context.Context, message *entity.Message, convID string) (*entity.Message, error)
	ListConversations(ctx context.Context, usr *entity.User) ([]*entity.Conversation, error)
	ListMessages(ctx context.Context, usr *entity.User, convID string) ([]*entity.Message, error)
}
