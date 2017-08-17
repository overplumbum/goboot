package main

import (
	"os"

	"api/commands"
	"api/config"

	"github.com/bugsnag/bugsnag-go"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func main() {
	gin.DisableConsoleColor()

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:              config.Config.BugSnagKey,
		PanicHandler:        func() {},
		ReleaseStage:        config.Config.ReleaseStage,
		NotifyReleaseStages: []string{config.STAGING, config.PRODUCTION},
	})
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:   "web",
			Usage:  "starts web server",
			Action: commands.Web,
		},
		{
			Name:   "migrate",
			Usage:  "create missing tables",
			Action: commands.Migrate,
		},
		{
			Name:   "schema",
			Usage:  "update schema.sql for registered models",
			Action: commands.Schema,
		},
		{
			Name:   "psql",
			Usage:  "run psql for dev db",
			Action: commands.Psql,
		},
	}

	app.Run(os.Args)
}
