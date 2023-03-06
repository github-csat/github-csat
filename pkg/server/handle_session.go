package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// example endpoint to demo reading/validating
// session cookie signature
func (h *Handlers) HandleSessionPing(c *gin.Context) {

	cookie, err := c.Cookie("github-session")
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	sessionDetails, err := h.validateAndParseCookie(cookie)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, sessionDetails)
}

type SessionDetails struct {
	GitHubUser string `json:"gitHubUser"`
}

// This is a stub method. Today the cookie is just a raw string w/ the
// GitHub username. Someday it will be a structured signed payload.
// We need to validate ths signature, then pull the SessionDetails
// out of the payload.
// This isn't being done client side because we want to use HTTPOnly cookies
func (h *Handlers) validateAndParseCookie(cookie string) (*SessionDetails, error) {
	// todo check the signature
	var cookieError error

	if cookieError != nil {
		return nil, errors.Wrap(cookieError, "Invalid Session Cookie")
	}

	// todo pull out the value post-validation
	gitHubLogin := cookie

	return &SessionDetails{
		GitHubUser: gitHubLogin,
	}, nil

}
