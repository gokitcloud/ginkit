package ginkit

func (e *Engine) Run(addr ...string) error {
	err := e.Router().Run(addr...)
	if err != nil {
		return err
	}

	return nil
}

func (e *Engine) RunSSLSelfSigned(addr string) error {
	err := GinRunSelfSignedSSL(e.Router(), addr)

	if err != nil {
		return err
	}

	return nil
}
