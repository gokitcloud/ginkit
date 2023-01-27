package ginkit

import (
	templateHTML "html/template"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/location"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

type (
	Engine struct {
		router          *gin.Engine
		version         string
		templateFuncMap templateHTML.FuncMap
	}
)

func NewDefault() (r *Engine) {
	e := New()

	e.Static("./web/assets")
	e.Templates("web/views", nil)

	e.Router().GET("/health", e.health)
	e.Router().GET("/ruok", e.health)
	e.Router().GET("/version", e.versionRoute)

	return e
}

func New() *Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(location.Default())
	r.Use(requestid.New())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// Use default logger
	r.Use(gin.Logger())

	// Cors Config - TODO Make CORS Configurarable
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin"}
	r.Use(cors.New(corsConfig))

	e := &Engine{
		router:          r,
		templateFuncMap: make(templateHTML.FuncMap),
	}

	// Copy base Template Functions
	for k, v := range globalHTMLFuncMap {
		e.templateFuncMap[k] = v
	}

	return e
}

func (e *Engine) Router() *gin.Engine {
	return e.router
}

func (e *Engine) SetVersion(version string) *Engine {
	e.version = version

	return e
}

func (e *Engine) versionRoute(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"version": e.version,
		},
	)
}

func (e *Engine) health(c *gin.Context) {
	ReturnData(c, http.StatusOK, gin.H{"status": "ok"})
}
