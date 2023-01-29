package ginkit

import (
	templateHTML "html/template"
	"net/http"
	"os"
	"time"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var (
	globalHTMLFuncMap = templateHTML.FuncMap{
		"env": func(key string) string {
			return os.Getenv(key)
		},
		"safeHTML": func(v string) templateHTML.HTML {
			return templateHTML.HTML(v)
		},
		"copy": func() string {
			return time.Now().Format("2006")
		},
		"time": func(t time.Time) string {
			return t.Format("2006-01-02 15:04:05")
		},
		"timeformat": func(t time.Time, format string) string {
			return t.Format(format)
		},
		"now": time.Now,
	}
)

func (e *Engine) Templates(views string, partials []string) *Engine {
	if views != "" {
		e.Router().HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
			Root:         views,
			Extension:    ".html",
			Master:       "layouts/main",
			Partials:     partials,
			Funcs:        e.templateFuncMap,
			DisableCache: true,
		})
	}

	return e
}

func (e *Engine) AddTemplateFunc(k string, v interface{}) *Engine {
	e.templateFuncMap[k] = v

	return e
}

func (e *Engine) HTML(template string, f any) func(*gin.Context) {
	return e.HTMLWithError(template, template, f)
}

func (e *Engine) HTMLWithError(template, errorTemplate string, f any) func(*gin.Context) {
	return func(c *gin.Context) {
		data, err := runTemplateFunc(c, f)
		if err != nil {
			c.HTML(http.StatusInternalServerError, errorTemplate, data)
		} else {
			c.HTML(http.StatusOK, template, data)
		}
	}
}

func runTemplateFunc(c *gin.Context, f any) (any, error) {
	var data any
	var err error

	switch f := f.(type) {
	case func(*gin.Context) (any, error):
		data, err = f(c)
	case func(*gin.Context):
		f(c)
	case func(Params) (any, error):
		data, err = f(NewParams(c.Params))
	case func() (any, error):
		data, err = f()
	case func() error:
		err = f()
	case nil:
	default:
		panic("Unknown function type")
	}

	return data, err
}
