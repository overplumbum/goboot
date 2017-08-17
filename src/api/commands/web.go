package commands

import (
	"log"
	"net/http"
	"os"

	"github.com/bugsnag/bugsnag-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"

	"api/config"
	"api/handlers"
	"api/models"
)

//noinspection GoUnusedParameter
func Web(c *cli.Context) error {
	if config.Config.DatabaseDSN != "" {
		log.Println("connecting to db")
		models.SetupDB()
		log.Println("connected")
	} else {
		log.Println("!!! WARNING !!! database not configured")
	}

	r := gin.Default()

	r.Use(bugsnaggin.AutoNotify())

	r.GET("/__lbheartbeat__", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "OK\n")
	})
	r.StaticFile("/favicon.ico", "assets/favicon.ico")
	r.StaticFile("/robots.txt", "assets/robots.txt")

	r.POST("/api/ping", handlers.PostPingHandler)

	authenticationInternal := gin.BasicAuth(gin.Accounts{
		"test": "test",
	})
	internal := r.Group("/", authenticationInternal)
	internal.GET("/__version__", handlers.GetVersionHandler)

	listen := os.Getenv("API_LISTEN")
	if listen == "" {
		listen = "127.0.0.1:8001"
	}
	log.Println("starting server at", listen)
	r.Run(listen)
	return nil
}
