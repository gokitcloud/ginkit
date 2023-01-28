package ginkit

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"net/url"
	"os"

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
func (e *Engine) SAMLGroup(path string, rootURL, entityID, metadataFile, certFile, keyFile string) *gin.RouterGroup {
	// TODO: Move to builtin Cert Store
	keyPair, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}

	// TODO - Support URL
	// idpMetadataURL, err := url.Parse("https://accounts.google.com/o/saml2/idp...")
	// if err != nil {
	// 	panic(err) // TODO handle error
	// }
	// idpMetadata, err := samlsp.FetchMetadata(context.Background(), http.DefaultClient,
	// 	*idpMetadataURL)
	// if err != nil {
	// 	panic(err) // TODO handle error
	// }
	xmlFile, err := os.ReadFile(metadataFile)
	if err != nil {
		panic(err) // TODO handle error
	}

	idpMetadata, err := samlsp.ParseMetadata(xmlFile)
	if err != nil {
		panic(err) // TODO handle error
	}

	rootURLParsed, err := url.Parse(rootURL)
	if err != nil {
		panic(err) // TODO handle error
	}

	samlSP, _ := samlsp.New(samlsp.Options{
		EntityID:           entityID,
		DefaultRedirectURI: "/",
		URL:                *rootURLParsed,
		AllowIDPInitiated:  true,
		Key:                keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:        keyPair.Leaf,
		IDPMetadata:        idpMetadata,
	})

	// TODO Does this need to be any?
	// e.Router().Any("/saml/*action", gin.WrapH(samlSP))
	e.Router().GET("/saml/*action", gin.WrapH(samlSP))
	e.Router().POST("/saml/*action", gin.WrapH(samlSP))

	restricted := e.Router().Group(path)
	restricted.Use(SAMLMiddleware(samlSP))

	return restricted
}

func SAMLMiddleware(samlSP *samlsp.Middleware) func(c *gin.Context) {
	return adapter.Wrap(samlSP.RequireAccount)
}
