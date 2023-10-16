package handler

import (
	"template-gin-api/internal/response"
	"template-gin-api/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Ctx struct {
	*gin.Context
	logger *zap.Logger
}

func (c *Ctx) Logger() *zap.Logger {
	return c.logger
}

func (c *Ctx) JSON(status int, data interface{}) {
	if status > 299 {
		switch v := data.(type) {
		case map[string]string:
			c.logger.WithOptions(zap.AddCallerSkip(1)).Error(v["error"])
		case gin.H:
			c.logger.WithOptions(zap.AddCallerSkip(1)).Error(v["error"].(string))
		case *response.ErrResponse:
			c.logger.WithOptions(zap.AddCallerSkip(1)).Error(v.Error)
		case error:
			c.logger.WithOptions(zap.AddCallerSkip(1)).Error(v.Error())
		}
	}
	c.Context.JSON(status, data)
}

type Handler func(*Ctx)

func New(handler Handler, logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&Ctx{
			c,
			logger.With(zap.String(utils.XRequestID, c.Request.Context().Value(utils.Key1).(string))),
		})
	}
}
