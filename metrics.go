package ginkit

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware(metric string, keys ...any) func(c *gin.Context) {
	return func(c *gin.Context) {
		metricValues := map[string]any{}

		for _, p := range keys {
			metricValue := parseContext(p, c)
			metricValues[fmt.Sprint(p)] = metricValue
		}

		// fmt.Println("# Metrics: ", metric, metricValues)
	}
}
