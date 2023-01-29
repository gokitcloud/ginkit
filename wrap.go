package ginkit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnData(c *gin.Context, status int, data any) {
	switch c.Query("_fmt") {
	case "json":
		c.JSON(status, data)
	case "toml":
		c.TOML(status, data)
	case "yaml":
		c.YAML(status, data)
	case "xml":
		c.XML(status, data)
	default:
		c.JSON(status, data)
	}
}

func WrapDataFunc(f func() (any, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		status := http.StatusOK

		data, err := f()
		if err != nil {
			status = http.StatusInternalServerError
			if data == nil {
				data = gin.H{
					"error": err.Error(),
				}
			}
		}

		ReturnData(c, status, data)
	}
}

func WrapDataFuncParams(f func(Params) (any, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		status := http.StatusOK

		data, err := f(NewParams(c.Params))
		if err != nil {
			status = http.StatusInternalServerError
			if data == nil {
				data = gin.H{
					"error": err.Error(),
				}
			}
		}

		ReturnData(c, status, data)
	}
}

func WrapDataFuncContext(f func(*gin.Context) (any, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		status := http.StatusOK

		data, err := f(c)
		if err != nil {
			status = http.StatusInternalServerError
			if data == nil {
				data = gin.H{
					"error": err.Error(),
				}
			}
		}

		ReturnData(c, status, data)
	}
}

func (e *Engine) Handle(httpMethod, relativePath string, funcs ...any) {
	handlers := []gin.HandlerFunc{}

	for _, f := range funcs {
		switch f := f.(type) {
		case func(*gin.Context) (any, error):
			handlers = append(handlers, WrapDataFuncContext(f))
		case func(*gin.Context):
			handlers = append(handlers, f)
		case func(Params) (any, error):
			handlers = append(handlers, WrapDataFuncParams(f))
		case func() (any, error):
			handlers = append(handlers, WrapDataFunc(f))
		default:
			panic("Unknown function type")
		}
	}

	e.Router().Handle(httpMethod, relativePath, handlers...)
}

func (e *Engine) GET(relativePath string, f any) {
	e.Handle(http.MethodGet, relativePath, f)
}

func (e *Engine) POST(relativePath string, f any) {
	e.Handle(http.MethodPost, relativePath, f)
}

func (e *Engine) PUT(relativePath string, f any) {
	e.Handle(http.MethodPut, relativePath, f)
}

func (e *Engine) HEAD(relativePath string, f any) {
	e.Handle(http.MethodHead, relativePath, f)
}

func (e *Engine) PATCH(relativePath string, f any) {
	e.Handle(http.MethodPatch, relativePath, f)
}

func (e *Engine) DELETE(relativePath string, f any) {
	e.Handle(http.MethodDelete, relativePath, f)
}

func (e *Engine) OPTIONS(relativePath string, f any) {
	e.Handle(http.MethodOptions, relativePath, f)
}
