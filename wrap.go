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
