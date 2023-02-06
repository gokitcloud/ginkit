package ginkit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e *Engine) SimpleTokenAuthGroup(path, token, header string) *RouterGroup {
	restricted := e.Router().Group(path)
	rg := RouterGroup{
		*restricted,
	}
	rg.Use(SimpleTokenAuthMiddleware(token, header))
	return &rg
}

func SimpleTokenAuthMiddleware(token, header string) func(c *gin.Context) {
	return func(c *gin.Context) {
		Authenticated := false
		if token == "" {
			Authenticated = true
		}
		for _, reqToken := range c.Request.Header[header] {
			if reqToken == token {
				Authenticated = true
			}
		}

		if !Authenticated {
			ReturnData(c, http.StatusForbidden, gin.H{"error": "must provide a valid token"})
			c.Abort()
			return
		}
	}
}

func SimpleTokenAuthOptionalMiddleware(token, header string) func(c *gin.Context) {
	return func(c *gin.Context) {
		Authenticated := false
		if token == "" {
			Authenticated = true
		}
		for _, reqToken := range c.Request.Header[header] {
			if reqToken == token {
				Authenticated = true
			}
		}

		c.Set("authenticated", Authenticated)
	}
}

func BasicAuthMiddleware(accounts gin.Accounts) func(c *gin.Context) {
	authMiddleware := gin.BasicAuth(accounts)
	return func(c *gin.Context) {
		Authenticated := c.GetBool("authenticated")

		if !Authenticated {
			authMiddleware(c)
		}
	}
}
