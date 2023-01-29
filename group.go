package ginkit

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	gin.RouterGroup
}

func (r *RouterGroup) HEAD(relativePath string, funcs ...any) {
	r.Handle(http.MethodHead, relativePath, funcs...)
}
func (r *RouterGroup) OPTIONS(relativePath string, funcs ...any) {
	r.Handle(http.MethodOptions, relativePath, funcs...)
}
func (r *RouterGroup) PATCH(relativePath string, funcs ...any) {
	r.Handle(http.MethodPatch, relativePath, funcs...)
}
func (r *RouterGroup) POST(relativePath string, funcs ...any) {
	r.Handle(http.MethodPost, relativePath, funcs...)
}
func (r *RouterGroup) PUT(relativePath string, funcs ...any) {
	r.Handle(http.MethodPut, relativePath, funcs...)
}
func (r *RouterGroup) DELETE(relativePath string, funcs ...any) {
	r.Handle(http.MethodDelete, relativePath, funcs...)
}
func (r *RouterGroup) GET(relativePath string, funcs ...any) {
	r.Handle(http.MethodGet, relativePath, funcs...)
}
func (r *RouterGroup) Handle(httpMethod, relativePath string, funcs ...any) {
	handlers := wrapHanders(funcs...)
	r.RouterGroup.Handle(httpMethod, relativePath, handlers...)
}
