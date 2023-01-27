package ginkit

import (
	templateHTML "html/template"
	"os"
	"time"

	gintemplate "github.com/foolin/gin-template"
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
