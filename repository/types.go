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
}

type Messages interface {
}
