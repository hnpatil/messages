package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/hnpatil/messages/entity"
	"github.com/sirupsen/logrus"
)

const (
	AuthCodeExpiredError int = 10001
	TokenInvalidError    int = 10002
	TokenExpiredError    int = 10003
)

type UsecaseError struct {
	ErrorCode int
	Message   string
}

func GetUsecaseError(err error) *UsecaseError {
	if err == nil {
		return nil
	}

	var e *UsecaseError
	if errors.As(err, &e) {
		return e
	}

	return nil
}

func (u *UsecaseError) Error() string {
	return u.Message
}

type Context interface {
	GetLogger() logrus.FieldLogger
	GetUserID() string
	GetContext() context.Context
}

type Token struct {
	AuthToken    string    `json:"auth_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type Auth interface {
	GenerateAuthCode(ctx Context, identifier string) error
	GetAuthToken(ctx Context, authCode string) (*Token, error)
	RefereshToken(ctx Context, refreshToken string) (*Token, error)
	Authenticate(ctx Context, token string) (string, error)
}

type Users interface {
	CreateUser(ctx Context, firstName string, lastName string) (*entity.User, error)
}

type Messages interface {
}
