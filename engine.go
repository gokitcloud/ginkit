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

func WrapH(d H) func(*gin.Context) {
	return func(c *gin.Context) {
		ReturnData(c, http.StatusOK, d)
	}
}

func WrapGinH(d gin.H) func(*gin.Context) {
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
	handlers := wrapHanders(funcs...)

	e.Router().NoRoute(handlers...)
}

func (e *Engine) Handle(httpMethod, relativePath string, funcs ...any) {
	handlers := wrapHanders(funcs...)

	if httpMethod == "Any" {
		for _, method := range anyMethods {
			e.Router().Handle(method, relativePath, handlers...)
		}
	} else {
		e.Router().Handle(httpMethod, relativePath, handlers...)
	}
}

func (e *Engine) Any(relativePath string, f ...any) {
	e.Handle("Any", relativePath, f...)
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

func (e *Engine) Redirect(location string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, location)
	}
}

func (e *Engine) NoRouteRedirect(location string) {
	e.Router().NoRoute(e.Redirect(location))
}

func (e *Engine) NoRouteHTML(template string, f any) {
	e.Router().NoRoute(e.HTML(template, f))
}
