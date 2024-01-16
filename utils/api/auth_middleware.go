package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/usecase"
)

func JwtAuthMiddleware(auth usecase.Auth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := extractToken(ctx)
		if bearerToken == "" {
			abortUnauthenticated(ctx)

			return
		}

		usrID, err := auth.Authenticate(usecase.NewContext(ctx), bearerToken)
		if err != nil {
			abortUnauthenticated(ctx)

			return
		}

		ctx.Set("identifier", usrID)
		ctx.Next()
	}
}

func abortUnauthenticated(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, &APIError{Error: "Unauthorized"})
	ctx.Abort()
}

func extractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}
