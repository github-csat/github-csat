package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rqlite/gorqlite"
)

type SubmitRequest struct {
	IssueURL string `json:"issueUrl"`
	Feedback string `json:"feedback"`
}

func (h *Handlers) APIHandleSubmit(c *gin.Context) {

	req := SubmitRequest{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(400, err)
		return
	}

	err := h.HandleSubmit(req)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, map[string]string{
		"status": "ok",
	})
}

func (h *Handlers) HandleSubmit(request SubmitRequest) error {
	ghUsername := "its_fake"
	issueURL := request.IssueURL
	feedback := request.Feedback

	conn, err := gorqlite.Open(h.RQLiteURL)
	if err != nil {
		return errors.Wrap(err, "connect to db")
	}

	_, err = conn.WriteOneParameterized(gorqlite.ParameterizedStatement{
		Query: `
		INSERT INTO satisfactions 
		(
			gh_username,
			issue_url,
			feedback
		) VALUES (?, ?, ?)`,
		Arguments: []interface{}{ghUsername, issueURL, feedback},
	})

	if err != nil {
		return errors.Wrap(err, "write rqlite")
	}

	return nil

}
