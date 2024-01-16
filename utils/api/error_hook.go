package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hnpatil/messages/entity"
	"github.com/hnpatil/messages/usecase"
	"github.com/loopfz/gadgeto/tonic"
)

type APIError struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

func tonicErrorHook(ctx *gin.Context, err error) (int, interface{}) {
	bindErr := &tonic.BindError{}
	if errors.As(err, bindErr) {
		return http.StatusBadRequest, &APIError{Error: bindErr.Error()}
	}

	if entity.IsNotFound(err) {
		return http.StatusNotFound, &APIError{Error: err.Error()}
	}

	if entity.IsConstraintError(err) {
		return http.StatusBadRequest, &APIError{Error: err.Error()}
	}

	if ue := usecase.GetUsecaseError(err); ue != nil {
		return http.StatusBadRequest, &APIError{Error: err.Error(), ErrorCode: ue.ErrorCode}
	}

	return http.StatusInternalServerError, &APIError{Error: err.Error()}
}
