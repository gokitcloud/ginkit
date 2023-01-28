package ginkit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (e *Engine) SimpleTokenAuthGroup(path, token, header string) *gin.RouterGroup {
	restricted := e.Router().Group(path)
	restricted.Use(SimpleTokenAuthMiddleware(token, header))

	return restricted
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
