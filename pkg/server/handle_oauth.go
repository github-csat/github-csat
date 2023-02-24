package server

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

var gitHubEndpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

var OauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
	Scopes:       []string{"read:user"},
	Endpoint:     gitHubEndpoint,
}

func (h *Handlers) HandleOauthRedirect(c *gin.Context) {
	log.Printf("redirect to: %s", OauthConfig.AuthCodeURL(""))
	c.Redirect(http.StatusFound, OauthConfig.AuthCodeURL(""))
}

func (h *Handlers) HandleOauthCallback(c *gin.Context) {
	val, ok := c.GetQuery("error")

	if ok {
		errDesc, _ := c.GetQuery("error_description")
		log.Printf("oauth err code: %s: %s\n", val, errDesc)
		// TODO: We can do better here on errors, maybe warrants a proper 500 page
		c.String(500, "There was an error during oauth flow. (Check the logs)")
		c.Abort()
		return
	}

	code, ok := c.GetQuery("code")

	token, err := OauthConfig.Exchange(oauth2.NoContext, code)

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

	c.String(200, "oauth flow was successful, what's shakin' %s (%s)", u.Name, u.Login)
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
		return nil, err
	}

	var u GitHubUser
	err = json.NewDecoder(resp.Body).Decode(&u)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
