package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func (h *Handlers) getOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     h.Config.GitHubClientID,
		ClientSecret: h.Config.GitHubClientSecret,
		Scopes:       []string{"read:user"},
		Endpoint:     h.Config.GitHubEndpoint,
	}
}

func (h *Handlers) HandleOauthRedirect(c *gin.Context) {
	log.Printf("redirect to: %s", h.getOauthConfig().AuthCodeURL(""))
	c.Redirect(http.StatusFound, h.getOauthConfig().AuthCodeURL(""))
}

func (h *Handlers) HandleOauthCallback(c *gin.Context) {
	val, ok := c.GetQuery("error")

	if ok {
		errDesc, _ := c.GetQuery("error_description")
		log.Printf("oauth err code: %s: %s\n", val, errDesc)
		// redirect and tell the frontend what to do with this
		c.Redirect(http.StatusFound, fmt.Sprintf("/auth/callback?error=%s&error_description=%s", val, errDesc))
		return
	}

	code, ok := c.GetQuery("code")

	token, err := h.getOauthConfig().Exchange(oauth2.NoContext, code)

	if err != nil {
		log.Println(err)
		c.String(500, "There was an error during the exchange process.")
		c.Abort()
		return
	}

	u, err := getUser(token.AccessToken)

	if err != nil {
		c.String(500, "There was an error retrieving this user")
	}

	h.setCookie(c, u.Login)
	c.Redirect(http.StatusFound, fmt.Sprintf("/auth/callback?name=%s&handle=%s", u.Name, u.Login))
}

func (h *Handlers) setCookie(c *gin.Context, login string) {
	// create a session, or sign a JWT that identifies this github user
	// for now fake this, just put the handle in the cookie w/ no sig
	sessionJWT := login
	sessionMaxAge := 5000 // seconds? minutes? who knows?
	sessionDomain := h.Config.SessionCookieDomain

	c.SetCookie(
		"github-session",
		sessionJWT,
		sessionMaxAge,
		"/api/",
		sessionDomain,
		true,
		true,
	)
}

type GitHubUser struct {
	AvatarUrl       string
	Company         string
	Id              int
	Login           string
	Name            string
	Location        string
	Blog            string
	Email           string
	Hireable        string
	Bio             string
	TwitterUsername string
	PublicRepos     int
	PublicGists     int
	Followers       int
	Following       int
}

func getUser(token string) (*GitHubUser, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)

	if err != nil {
		return nil, errors.Wrap(err, "create request")
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.Wrap(err, "send user request with oath token")
	}

	var u GitHubUser
	err = json.NewDecoder(resp.Body).Decode(&u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
