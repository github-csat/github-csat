package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
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
	router.GET("/submit", handlers.HandleSubmit)
	err = router.Run(conf.GinAddress)

	return errors.Wrap(err, "run gin http server")
}

func (h *Handlers) HandleSubmit(c *gin.Context) {
	issueURL := c.Query("issue")
	responseValue := c.Query("response")

	c.String(http.StatusOK, "Your response on %s has been recorded as %s", issueURL, responseValue)
}
