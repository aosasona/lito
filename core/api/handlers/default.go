package handlers

import (
	e "errors"

	"go.trulyao.dev/lito/pkg/errors"
	"go.trulyao.dev/lito/pkg/logger"
	"go.trulyao.dev/lito/pkg/types"
)

func ErrorHandler(ctx *types.Context, err error) {
	ctx.Logger.Error("Failed to handle request",
		logger.Param{Key: "domain", Value: ctx.Request.Host},
		logger.Param{Key: "path", Value: ctx.Request.URL.Path},
		logger.Param{Key: "error", Value: err.Error()},
	)

	var (
		code          = 500
		reportedError = e.New("An error occurred while handling this request, see logs for details if you own this application")
	)
	if pe, ok := err.(*errors.ProxyError); ok {
		code = pe.Code()
		reportedError = pe
	}

	ctx.Error(code, reportedError)
}
