package middleware

import (
	"bytes"
	"context"
	"io"
	"template-gin-api/config"
	"template-gin-api/utils"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type middleware struct {
	cfg    *config.Config
	logger *zap.Logger
}

func NewMiddleware(cfg *config.Config, logger *zap.Logger) *middleware {
	return &middleware{
		cfg:    cfg,
		logger: logger,
	}
}

func (m *middleware) BasicAuthenication() gin.HandlerFunc {
	return gin.BasicAuthForRealm(gin.Accounts{
		"foo": "bar",
	}, "Restricted")
}

func (m *middleware) CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  []string{m.cfg.App.Cors.Origin},
		AllowMethods:  []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Origin", "Accept"},
		ExposeHeaders: []string{"Content-Length"},
	})
}

func (m *middleware) JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetAccepted(utils.ApplicationJSON)
		c.Next()
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (m *middleware) LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		xid := c.Request.Header.Get(utils.XRequestID)
		if xid == "" {
			xid = uuid.NewString()
			c.Request.Header.Add(utils.XRequestID, xid)
		}

		logger := m.logger.With(zap.String(utils.XRequestID, xid))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), utils.Key1, xid))

		// https://github.com/gin-gonic/gin/issues/1363
		// https://github.com/gin-gonic/gin/issues/961
		var reqBuf bytes.Buffer
		reqTee := io.TeeReader(c.Request.Body, &reqBuf)
		reqBody, _ := io.ReadAll(reqTee)
		c.Request.Body = io.NopCloser(&reqBuf)

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		logger.Debug(utils.RequestInfoMsg,
			zap.String("method", c.Request.Method),
			zap.String("host", c.Request.Host),
			zap.String("path_uri", c.Request.RequestURI),
			zap.String("url.host", c.Request.URL.Host),
			zap.String("url.path", c.Request.URL.Path),
			zap.String("url.raw-path", c.Request.URL.RawPath),
			zap.String("remote_addr", c.Request.RemoteAddr),
			zap.String("body", string(reqBody)),
		)

		c.Next()

		logger.Debug(utils.ResponseInfoMsg,
			zap.String("body", w.body.String()),
		)

		logger.Info(utils.SummaryInfoMsg,
			zap.String("method", c.Request.Method),
			zap.String("path_uri", c.Request.RequestURI),
			zap.Duration("duration", time.Since(start)),
			zap.Int("status_code", c.Writer.Status()),
		)
	}
}
