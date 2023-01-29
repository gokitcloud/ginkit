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

func WrapErrorFunc(f func() error) func(*gin.Context) {
	return func(c *gin.Context) {
		status := http.StatusOK
		var data any

		err := f()
		if err != nil {
			status = http.StatusInternalServerError
			data = gin.H{
				"error": err.Error(),
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

func WrapH(d gin.H) func(*gin.Context) {
	return func(c *gin.Context) {
		ReturnData(c, http.StatusOK, d)
	}
}

func WrapBytes(b []byte) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", b)
	}
}

func WrapString(s string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte(s))
	}
}

func (e *Engine) NoRoute(funcs ...any) {
	handlers := e.wrapHanders(funcs...)

	e.Router().NoRoute(handlers...)
}

func (e *Engine) Handle(httpMethod, relativePath string, funcs ...any) {
	handlers := e.wrapHanders(funcs...)
	e.Router().Handle(httpMethod, relativePath, handlers...)
}

func (e *Engine) wrapHanders(funcs ...any) []gin.HandlerFunc {
	handlers := []gin.HandlerFunc{}

	for _, f := range funcs {
		switch f := f.(type) {
		case gin.HandlersChain:
			handlers = append(handlers, f...)
		case gin.HandlerFunc:
			handlers = append(handlers, f)
		case http.HandlerFunc:
			handlers = append(handlers, gin.WrapF(f))
		case http.Handler:
			handlers = append(handlers, gin.WrapH(f))
		case func(*gin.Context) (any, error):
			handlers = append(handlers, WrapDataFuncContext(f))
		case func(Params) (any, error):
			handlers = append(handlers, WrapDataFuncParams(f))
		case func() (any, error):
			handlers = append(handlers, WrapDataFunc(f))
		case func() error:
			handlers = append(handlers, WrapErrorFunc(f))
		case string:
			handlers = append(handlers, WrapString(f))
		case []byte:
			handlers = append(handlers, WrapBytes(f))
		case gin.H:
			handlers = append(handlers, WrapH(f))
		default:
			panic("Unknown function type")
		}
	}

	return handlers
}

func (e *Engine) GET(relativePath string, f ...any) {
	e.Handle(http.MethodGet, relativePath, f...)
}

func (e *Engine) POST(relativePath string, f ...any) {
	e.Handle(http.MethodPost, relativePath, f...)
}

func (e *Engine) PUT(relativePath string, f ...any) {
	e.Handle(http.MethodPut, relativePath, f...)
}

func (e *Engine) HEAD(relativePath string, f ...any) {
	e.Handle(http.MethodHead, relativePath, f...)
}

func (e *Engine) PATCH(relativePath string, f ...any) {
	e.Handle(http.MethodPatch, relativePath, f...)
}

func (e *Engine) DELETE(relativePath string, f ...any) {
	e.Handle(http.MethodDelete, relativePath, f...)
}

func (e *Engine) OPTIONS(relativePath string, f ...any) {
	e.Handle(http.MethodOptions, relativePath, f...)
}
