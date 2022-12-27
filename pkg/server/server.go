package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	RQLiteURL string
}

func Main() error {

	conf, err := LoadConfig()
	if err != nil {
		return errors.Wrap(err, "load config")
	}

	handlers := &Handlers{
		RQLiteURL: conf.RQLiteURL,
	}

	fmt.Println(handlers.RQLiteURL)

	router := gin.Default()

	webRoutes := router.Group("/")
	webRoutes.GET("/", handleIndex)
	webRoutes.GET("/submit", handlers.HandleSubmit)
	webRoutes.GET("/satisfactions", handlers.HandleSatisfactions)

	apiRoutes := router.Group("/api")
	apiRoutes.POST("/submit", handlers.APIHandleSubmit)

	err = router.Run(conf.GinAddress)

	return errors.Wrap(err, "run gin http server")
}

type Satisfaction struct {
	GhUsername   string
	IssueUrl     string
	Feedback     string
	SatisfiedAt  *time.Time
	IssueCreated *time.Time
	IssueClosed  *time.Time
}

func (h *Handlers) HandleSatisfactions(c *gin.Context) {
	body := []string{
		"SELECT * FROM satisfactions",
	}

	jsonBody, err := json.Marshal(body)

	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/db/query", h.RQLiteURL),
		"application/json",
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(responseBytes)
}

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
	// todo something with the request
}

func (h *Handlers) HandleSubmit(c *gin.Context) {
	ghUsername := "its_fake"
	issueURL := c.Query("issue_url")
	feedback := c.Query("feedback")

	body := [][]string{
		{
			`
		INSERT INTO satisfactions 
		(
			gh_username,
			issue_url,
			feedback
		) VALUES (?, ?, ?)`,
			ghUsername, issueURL, feedback,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/db/execute?pretty&timings", h.RQLiteURL),
		"application/json",
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	defer resp.Body.Close()

	responseBytes, err := io.ReadAll(resp.Body)

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(fmt.Sprintf(
		`
<html><body>
  We sent 
    <p><pre>%s</pre><p> 
  and got 
    <p><pre>%s</pre><p>
</body></html>`, jsonBody, responseBytes)))
}

func handleIndex(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(index))
}
