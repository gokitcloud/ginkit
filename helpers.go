package ginkit

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func parseContext(p any, c *gin.Context) any {
	var paramValue any
	switch p.(type) {
	case string:
		param := strings.Split(p.(string), ":")
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
		case "param":
			paramValue = c.Param(param[1])
		case "header":
			paramValue = c.Request.Header.Get(param[1])
		case "session":
			session := sessions.Default(c)
			paramValue = session.Get(param[1])
		}
	case func(c *gin.Context) any:
		paramValue = p.(func(c *gin.Context) any)(c)
	}

	return paramValue
}
