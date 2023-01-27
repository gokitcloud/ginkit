package ginkit

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func (e *Engine) RBACPathGroup(path, model, policy, tokenHeader string) *gin.RouterGroup {
	restricted := e.Router().Group(path)
	restricted.Use(RBACPathMiddleware(model, policy, tokenHeader))

	return restricted
}

func RBACPathMiddleware(model, policy, tokenHeader string) func(c *gin.Context) {
	e, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		allowed := false
		method := c.Request.Method
		path := c.Request.URL.Path

		for _, reqToken := range c.Request.Header[tokenHeader] {
			allowed, err = e.Enforce(reqToken, path, method)
			if err != nil {
				c.JSON(
					http.StatusForbidden,
					gin.H{"error": err.Error()},
				)
				c.Abort()
				return
			}
			if allowed {
				continue
			}
		}

		if !allowed {
			c.JSON(
				http.StatusForbidden,
				gin.H{"error": "must provide a valid token"},
			)
			c.Abort()
			return
		}
	}
}

func (e *Engine) RBACParamGroup(path, model, policy, tokenHeader string, params []string) *gin.RouterGroup {
	restricted := e.Router().Group(path)
	restricted.Use(RBACParamMiddleware(model, policy, tokenHeader, params))

	return restricted
}

func RBACParamMiddleware(model, policy, tokenHeader string, params []string) func(c *gin.Context) {
	e, err := casbin.NewEnforcer(model, policy)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		allowed := false
		var paramValues []string

		for _, p := range params {
			paramValues = append(paramValues, c.Param(p))
		}

		for _, reqToken := range c.Request.Header[tokenHeader] {
			args := []interface{}{reqToken}
			for _, v := range paramValues {
				args = append(args, v)
			}
			log.Println(args)
			allowed, err = e.Enforce(args...)
			if err != nil {
				c.JSON(
					http.StatusForbidden,
					gin.H{"error": err.Error()},
				)
				c.Abort()
				return
			}
			if allowed {
				continue
			}
		}

		if !allowed {
			c.JSON(
				http.StatusForbidden,
				gin.H{"error": "must provide a valid token"},
			)
			c.Abort()
			return
		}
	}
}
