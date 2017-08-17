package commands

import (
	"log"

	"github.com/urfave/cli"

	"api/models"
)

//noinspection GoUnusedParameter
func Migrate(c *cli.Context) error {
	log.Println("connecting to db")
	models.SetupDB()
	log.Println("creating missing tables")
	log.Println("this kind of migration won't modify any existing tables")
	models.Migrate()
	return nil
}
