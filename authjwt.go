package ginkit

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/pkg/errors"
)

func (e *Engine) JWTAuthGroup(path, jwksURL, boundIssuer, jwtheader, claimid string) *RouterGroup {
	restricted := e.Router().Group(path)
	rg := RouterGroup{
		*restricted,
	}
	rg.Use(JWTAuthMiddleware(jwksURL, boundIssuer, jwtheader, claimid))
	return &rg
}

func JWTAuthMiddleware(jwksURL, boundIssuer, jwtheader, claimid string) func(c *gin.Context) {
	return func(c *gin.Context) {
		Authenticated := c.GetBool("authenticated")

		if !Authenticated {
			return
		}
		token := c.Request.Header.Get(jwtheader)

		keySet, err := getKeySet(jwksURL)
		if err != nil {
			log.Println(err)
			return
		}

		claims, err := validateJWT(keySet, boundIssuer, token)
		if err != nil {
			return
		}

		c.Set("authenticated", true)
		c.Set("claims", claims)
		c.Set("user", claims[claimid])
	}
}

func JWTAuthCachedMiddleware(jwksURL, boundIssuer, jwtheader, claimid string) func(c *gin.Context) {
	keySet, err := getKeySet(jwksURL)
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		Authenticated := c.GetBool("authenticated")

		if !Authenticated {
			return
		}

		token := c.Request.Header.Get(jwtheader)

		claims, err := validateJWT(keySet, boundIssuer, token)
		if err != nil {
			return
		}

		c.Set("authenticated", true)
		c.Set("claims", claims)
		c.Set("user", claims[claimid])
	}
}

func getKeySet(jwksURL string) (jwk.Set, error) {
	keySet, err := jwk.Fetch(context.Background(), jwksURL)
	if err != nil {
		return nil, errors.Errorf("failed to parse JWK from %s: %v", jwksURL, err)
	}
	if keySet.Len() < 1 {
		return nil, errors.Errorf("%s did not return any key", jwksURL)
	}

	return keySet, nil
}

func validateJWT(keySet jwk.Set, boundIssuer, tokenString string) (map[string]any, error) {
	token, err := jwt.ParseString(tokenString, jwt.WithIssuer(boundIssuer), jwt.WithKeySet(keySet))
	if err != nil {
		return nil, errors.Errorf("failed to validate the jwt: %v", err)
	}
	claims := token.PrivateClaims()

	return claims, nil
}
