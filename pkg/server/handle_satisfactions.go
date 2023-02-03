package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/rqlite/gorqlite"
	"time"
)

type Satisfaction struct {
	ID             int        `json:"id"`
	GitHubUsername string     `json:"gitHubUsername"`
	IssueUrl       string     `json:"issueUrl"`
	Feedback       string     `json:"feedback"`
	SatisfiedAt    time.Time  `json:"satisfiedAt"`
	IssueCreated   *time.Time `json:"issueCreated"`
	IssueClosed    *time.Time `json:"issueClosed"`
}

func (h *Handlers) APIHandleSatisfactions(c *gin.Context) {
	satisfactions, err := h.HandleSatisfactions()

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, satisfactions)
}
func (h *Handlers) HandleSatisfactions() ([]Satisfaction, error) {

	conn, err := gorqlite.Open(h.RQLiteURL)
	if err != nil {
		return nil, errors.Wrap(err, "connect to db")
	}

	results, err := conn.QueryOne(`
SELECT 
    id,
    gh_username,
    issue_url,
    feedback,
    satisfied_at,
    issue_created,
    issue_closed
FROM satisfactions
`)

	if err != nil {
		return nil, errors.Wrap(err, "query rqlite")
	}

	var satisfactions []Satisfaction
	for results.Next() {
		var satisfaction Satisfaction
		var issueCreated, issueClosed gorqlite.NullTime
		err := results.Scan(
			&satisfaction.ID,
			&satisfaction.GitHubUsername,
			&satisfaction.IssueUrl,
			&satisfaction.Feedback,
			&satisfaction.SatisfiedAt,
			&issueCreated,
			&issueClosed,
		)
		if err != nil {
			return nil, errors.Wrap(err, "scan row")
		}
		if issueCreated.Valid {
			satisfaction.IssueCreated = &issueCreated.Time
		}
		if issueClosed.Valid {
			satisfaction.IssueCreated = &issueClosed.Time
		}
		satisfactions = append(satisfactions, satisfaction)
	}

	return satisfactions, nil
}
