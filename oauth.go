package ginkit

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (e *Engine) OAuthGroup(path string, authServerURL string, config oauth2.Config) *gin.RouterGroup {
	if !e.sessionsEnabled {
		log.Fatal("Must enable sessions.")
	}
	noauth := e.Router().Group("/")
	noauth.GET("/oauth2", oauthEndpoint(config))
	noauth.GET("/logout", LogoutRoute)

	authorized := e.Router().Group(path)
	authorized.Use(oauthMiddleware(authServerURL, config))

	return authorized
}

func (e *Engine) OAuthMW(path string, authServerURL string, config oauth2.Config) gin.HandlerFunc {
	if !e.sessionsEnabled {
		log.Fatal("Must enable sessions.")
	}
	noauth := e.Router().Group("/")
	noauth.GET("/oauth2", oauthEndpoint(config))
	noauth.GET("/logout", LogoutRoute)

	return oauthMiddleware(authServerURL, config)
}

func genCodeChallengeS256(s string) string {
	s256 := sha256.Sum256([]byte(s))
	return base64.URLEncoding.EncodeToString(s256[:])
}

// TODO: Add errorFunc
func oauthMiddleware(authServerURL string, config oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("token") == nil {
			// TODO: Param for s256example
			u := config.AuthCodeURL(c.Request.URL.Path,
				oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256("s256example")),
				oauth2.SetAuthURLParam("code_challenge_method", "S256"))

			c.Redirect(http.StatusFound, u)
			c.Abort()
			err := session.Save()
			if err != nil {
				// TODO: replace with errorFunc
				c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			return
		}

		if session.Get("polled") == nil {
			// TODO: replace with config
			resp, err := http.Get(fmt.Sprintf("%s/test?access_token=%s", authServerURL, session.Get("token")))
			if err != nil {
				// TODO: replace with errorFunc
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()
				return

			}
			defer resp.Body.Close()

			// TODO: was return a 200?

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				// TODO: replace with errorFunc
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				c.Abort()

				return
			}
			e := map[string]interface{}{}
			err = json.Unmarshal(body, &e)
			if err != nil {
				log.Println(err)
				// TODO: replace with errorFunc

				u := config.AuthCodeURL(c.Request.URL.Path,
					oauth2.SetAuthURLParam("code_challenge", genCodeChallengeS256("s256example")),
					oauth2.SetAuthURLParam("code_challenge_method", "S256"))

				c.Redirect(http.StatusFound, u)
				c.Abort()
				err := session.Save()
				if err != nil {
					// TODO: replace with errorFunc
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					c.Abort()
					return
				}

				return
			}

			for k, v := range e {
				session.Set("oauth_"+k, v)
			}

			session.Set("polled", true)

			err = session.Save()
			if err != nil {
				// TODO: replace with errorFunc
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func oauthEndpoint(config oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)

		err := c.Request.ParseForm()
		if err != nil {
			// TODO: replace with errorFunc
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		state := c.Request.Form.Get("state")
		code := c.Request.Form.Get("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
			return
		}
		token, err := config.Exchange(context.Background(), code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		session.Set("token", token.AccessToken)
		err = session.Save()
		if err != nil {
			// TODO: replace with errorFunc
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if state == "" {
			c.JSON(200, gin.H{"state": state, "token": token})
			return
		} else {
			c.Redirect(http.StatusFound, state)
			return
		}
	}
}
