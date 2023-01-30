package ginkit

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func parseContext(p any, c *gin.Context) any {
	var paramValue any
	switch p := p.(type) {
	case string:
		param := strings.Split(p, ":")
		if len(param) < 2 {
			ReturnData(c, http.StatusForbidden, gin.H{"error": "invalid RBAC config"})
			c.Abort()
			return nil
		}

		switch param[0] {
		case "request":
			switch param[1] {
			case "method":
				paramValue = c.Request.Method
			case "path":
				paramValue = c.Request.URL.Path
			}
		case "value":
			paramValue = param[1]
		case "context":
			paramValue, _ = c.Get(param[1])
		case "param":
			paramValue = c.Param(param[1])
		case "header":
			paramValue = c.Request.Header.Get(param[1])
		case "session":
			session := sessions.Default(c)
			paramValue = session.Get(param[1])
		}
	case func(c *gin.Context) any:
		paramValue = p(c)
	}

	return paramValue
}

func wrapHanders(funcs ...any) []gin.HandlerFunc {
	handlers := []gin.HandlerFunc{}

	for _, f := range funcs {
		switch f := f.(type) {
		case gin.HandlersChain:
			handlers = append(handlers, f...)
		case []gin.HandlerFunc:
			handlers = append(handlers, f...)
		case gin.HandlerFunc:
			handlers = append(handlers, f)
		case func(*gin.Context):
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
			handlers = append(handlers, WrapGinH(f))
		case H:
			handlers = append(handlers, WrapH(f))
		case map[string]any:
			handlers = append(handlers, WrapH(f))
		default:
			panic("Unknown function type")
		}
	}

	return handlers
}
