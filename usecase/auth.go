package usecase

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hnpatil/messages/repository"
	"github.com/hnpatil/messages/utils/config"
	"golang.org/x/oauth2"
)

type authImpl struct {
	authCodes  repository.AuthCodes
	authSecret []byte
}

func NewAuth(authCodes repository.AuthCodes, cfg *config.Config) Auth {
	return &authImpl{
		authCodes:  authCodes,
		authSecret: []byte(cfg.GetValue(config.AUTH_SECRET)),
	}
}

func (a *authImpl) GenerateAuthCode(ctx Context, identifier string) error {
	code := oauth2.GenerateVerifier()
	expiresAt := time.Now().Add(10 * time.Minute)

	authCode, err := a.authCodes.CreateAuthCode(ctx.GetContext(), identifier, code, &expiresAt)
	if err != nil {
		return err
	}

	ctx.GetLogger().Infof("Sending auth code %s to user %s", authCode.Code, identifier)

	return nil
}

func (a *authImpl) GetAuthToken(ctx Context, authCode string) (*Token, error) {
	mdl, err := a.authCodes.GetAuthCode(ctx.GetContext(), authCode)
	if err != nil {
		return nil, err
	}

	if mdl.ExpiresAt.Before(time.Now()) {
		return nil, &UsecaseError{
			ErrorCode: AuthCodeExpiredError,
			Message:   "expired auth code",
		}
	}

	return a.getToken(mdl.Identifier)
}

func (a *authImpl) getToken(identifier string) (*Token, error) {
	authExpiry := time.Now().Add(time.Hour)

	authToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"identifier": identifier,
		"exp":        authExpiry.Unix(),
	})

	authString, err := authToken.SignedString(a.authSecret)
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": false,
		"identifier": identifier,
		"exp":        time.Now().Add(30 * 24 * time.Hour).Unix(),
	})

	refreshString, err := refreshToken.SignedString(a.authSecret)
	if err != nil {
		return nil, err
	}

	return &Token{
		AuthToken:    authString,
		RefreshToken: refreshString,
		ExpiresAt:    authExpiry,
	}, nil
}

func (a *authImpl) RefereshToken(ctx Context, refreshToken string) (*Token, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return a.authSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, &UsecaseError{ErrorCode: TokenInvalidError, Message: "Invalid token"}
	}

	return a.getToken(fmt.Sprintf("%s", claims["identifier"]))
}

func (a *authImpl) Authenticate(ctx Context, authToken string) (string, error) {
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return a.authSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", &UsecaseError{ErrorCode: TokenInvalidError, Message: "Invalid token"}
	}

	if authorized, ok := claims["authorized"].(bool); !ok || !authorized {
		return "", &UsecaseError{ErrorCode: TokenInvalidError, Message: "Invalid token"}
	}

	return fmt.Sprintf("%s", claims["identifier"]), nil
}
