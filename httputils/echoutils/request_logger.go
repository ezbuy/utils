package echoutils

import (
	"net/http"

	"github.com/apex/log"
	"github.com/labstack/echo"
)

const (
	CtxLoggerKey     = "_logger"
	ReqIdHeaderField = "X-Req-Id"
)

type ReqLogger struct {
	r      *http.Request
	logger *log.Entry
}

func RequestLogger(c *echo.Context) *log.Entry {
	req := c.Request()
	if rLogger, ok := c.Get(CtxLoggerKey).(*ReqLogger); ok && rLogger.r == req {
		// same request
		return rLogger.logger
	}

	reqId := GenReqId()
	c.Response().Header().Set(ReqIdHeaderField, reqId)

	rLogger := &ReqLogger{
		r:      req,
		logger: logger.WithField("ReqId", reqId),
	}
	c.Set(CtxLoggerKey, rLogger)

	return rLogger.logger
}
