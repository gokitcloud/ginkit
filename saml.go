package ginkit

import (
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"net/url"
	"os"

	"github.com/crewjam/saml"
	"github.com/crewjam/saml/samlsp"
	"github.com/gin-gonic/gin"
	adapter "github.com/gwatts/gin-adapter"
)

// metadataFile metadata.xml
// certFile myservice.cert
// keyFile myservice.key
// entityID asdfasdf
// rootURL https://devlocal.site:8000
// ACS URL: /saml/acs
// TODO: Implement SAMLGroupConfig
type (
	SAMLGroupConfig struct {
		MetaDataFile string
		MetaDataURL  string
		CertFile     string
		KeyFile      string
		EntityID     string
		RootURL      string
		ParamMap     map[string]string
		AllowSkip    bool
	}
)

func (e *Engine) SAMLGroup(path string, config *SAMLGroupConfig) *gin.RouterGroup {
	samlSP := e.SAMLInit(config)
	restricted := e.Router().Group(path)
	restricted.Use(SAMLMiddleware(samlSP))
	restricted.Use(SAMLtoParamsMapMiddleware(config.ParamMap))

	return restricted
}

func (e *Engine) SAMLInit(config *SAMLGroupConfig) *samlsp.Middleware {
	// TODO: Move to builtin Cert Store
	keyPair, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	var idpMetadata *saml.EntityDescriptor
	// TODO - Support URL
	if config.MetaDataURL != "" {
		idpMetadataURL, err := url.Parse(config.MetaDataURL)
		if err != nil {
			panic(err) // TODO handle error
		}
		idpMetadata, err = samlsp.FetchMetadata(context.Background(), http.DefaultClient, *idpMetadataURL)
		if err != nil {
			panic(err) // TODO handle error
		}
	} else {
		xmlFile, err := os.ReadFile(config.MetaDataFile)
		if err != nil {
			panic(err) // TODO handle error
		}
		idpMetadata, err = samlsp.ParseMetadata(xmlFile)
		if err != nil {
			panic(err) // TODO handle error
		}

	}

	rootURLParsed, err := url.Parse(config.RootURL)
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		EntityID:           config.EntityID,
		DefaultRedirectURI: "/",
		URL:                *rootURLParsed,
		AllowIDPInitiated:  true,
		Key:                keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:        keyPair.Leaf,
		IDPMetadata:        idpMetadata,
	})

	// TODO Does this need to be any?
	// TODO make /saml path configurable
	// e.Router().Any("/saml/*action", gin.WrapH(samlSP))
	e.Router().GET("/saml/*action", gin.WrapH(samlSP))
	e.Router().POST("/saml/*action", gin.WrapH(samlSP))

	return samlSP
}

func SAMLMiddleware(samlSP *samlsp.Middleware) func(c *gin.Context) {
	return func(c *gin.Context) {
		Authenticated := c.GetBool("authenticated")

		if Authenticated {
			return
		}

		adapter.Wrap(samlSP.RequireAccount)(c)
	}
}

func SAMLtoParamsMiddleware(params ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Range over all params which we want to add from the saml context
		for _, param := range params {
			// Get the saml attribute from the saml context
			value := samlsp.AttributeFromContext(c.Request.Context(), param)

			// Determine if there was already a param
			_, exists := c.Params.Get(param)

			// If the param already exists overwrite else append param
			if exists {
				for i, entry := range c.Params {
					if entry.Key == param {
						c.Params[i] = gin.Param{
							Key:   param,
							Value: value,
						}
						return
					}
				}
			} else {
				c.Params = append(c.Params, gin.Param{
					Key:   param,
					Value: value,
				})
			}
		}
	}
}

func SAMLtoParamsMapMiddleware(params map[string]string) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Range over all params which we want to add from the saml context
		for param, key := range params {
			// Get the saml attribute from the saml context
			value := samlsp.AttributeFromContext(c.Request.Context(), key)

			// Determine if there was already a param
			_, exists := c.Params.Get(param)

			// If the param already exists overwrite else append param
			if exists {
				for i, entry := range c.Params {
					if entry.Key == param {
						c.Params[i] = gin.Param{
							Key:   param,
							Value: value,
						}
						return
					}
				}
			} else {
				c.Params = append(c.Params, gin.Param{
					Key:   param,
					Value: value,
				})
			}
		}
	}
}
