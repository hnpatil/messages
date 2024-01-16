package usecase

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ucContext struct {
	context.Context
}

func NewContext(ctx context.Context) Context {
	return &ucContext{Context: ctx}
}

func (c *ucContext) GetLogger() logrus.FieldLogger {
	lg := c.Value("logger")
	if lg == nil {
		logrus.Fatal("no logger set")
	}

	return lg.(logrus.FieldLogger)
}

func (c *ucContext) GetUserID() string {
	usr := c.Value("identifier")
	if usr == nil {
		return ""
	}

	return usr.(string)
}

func (c *ucContext) GetContext() context.Context {
	return c.Context
}
