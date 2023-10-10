package types

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.trulyao.dev/lito/pkg/logger"
)

type ErrorHandler func(ctx *Context, err error)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Logger   logger.Logger
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
	}
}

func (c *Context) Status(code int) {
	c.Response.WriteHeader(code)
}

func (c *Context) Header(key, value string) *Context {
	c.Response.Header().Set(key, value)
	return c
}

func (c *Context) Send(data []byte) {
	c.Response.Write(data)
}

// JSON sends a JSON response with status code
//
// Warning: using this method automatically sets the header and status code, do not use c.Status() or c.Header(k,v) before calling this method
func (c *Context) JSON(code int, obj interface{}) {
	c.Header("Content-Type", "application/json")

	data, err := json.Marshal(obj)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		c.Send([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
		return
	}

	c.Status(code)
	c.Send(data)
}

func (c *Context) Error(code int, err error) {
	c.Status(code)
	c.Send([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
