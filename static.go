package ginkit

func (e *Engine) Static(assets string) (*Engine) {
	e.Router().Static("/assets", assets)
	e.Router().Static("/css", assets+"/css")
	e.Router().Static("/font", assets+"/font")
	e.Router().Static("/img", assets+"/img")
	e.Router().Static("/js", assets+"/js")

	e.Router().StaticFile("/favicon.ico", assets+"/favicon.ico")

	return e
}
