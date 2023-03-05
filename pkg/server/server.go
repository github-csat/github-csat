package server

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Handlers struct {
	Config
}

func Main() error {

	conf, err := LoadConfig()
	if err != nil {
		return errors.Wrap(err, "load config")
	}

	handlers := &Handlers{
		Config: *conf,
	}

	fmt.Println(handlers.RQLiteURL)

	router := gin.Default()

	apiRoutes := router.Group("/api")
	apiRoutes.POST("/submit", handlers.APIHandleSubmit)
	apiRoutes.GET("/satisfactions", handlers.APIHandleSatisfactions)

	apiRoutes.GET("oauth/redirect", handlers.HandleOauthRedirect)
	apiRoutes.GET("oauth/callback", handlers.HandleOauthCallback)

	if conf.StaticDir != "" {
		fmt.Println("adding static handler")
		router.Use(static.Serve("/", static.LocalFile(conf.StaticDir, true)))
	} else if conf.ProxyFrontendURL != nil {
		setUpDevProxy(handlers, router)
	} else {
		return errors.New("need at least one of STATIC_DIR or PROXY_FRONTEND to serve web assets")
	}

	err = router.Run(conf.GinAddress)

	return errors.Wrap(err, "run gin http server")
}
