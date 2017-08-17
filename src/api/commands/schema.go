package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"api/dry"
	"api/models"

	"github.com/jinzhu/gorm"
	"github.com/urfave/cli"
)

//noinspection GoUnusedParameter
func Schema(c *cli.Context) error {
	//models.SetupDB()

	pg := NewPostgreSQL()
	defer pg.Remove()

	time.Sleep(time.Second * 5)
	{
		db, err := gorm.Open("postgres", pg.DSN())
		dry.Check(err)

		db.SingularTable(true)
		db.LogMode(true)

		models.DB = db
	}

	log.Println("creating missing tables")
	log.Println("this kind of migration won't modify any existing tables")
	models.Migrate()

	dump := pg.Dump(models.Models)
	dry.Check(ioutil.WriteFile("schema.sql", dump, 0666))

	log.Println("schema.sql dump complete")
	return nil
}

type PostgreSQL struct {
	Password      string
	ContainerName string
	Database      string
	Port          string
}

func NewPostgreSQL() *PostgreSQL {
	log.Println("starting temp postgresql instance via docker")
	p := PostgreSQL{
		Password:      "jee7cee8Fi",
		Database:      "gorm",
		ContainerName: "gorm",
	}

	// don't check status here, missing container is ok
	exec.Command("docker", "container", "rm", "-f", p.ContainerName).Run()

	args := []string{
		"run",
		"--rm",
		"--name=" + p.ContainerName,
		"-d",
		"-e", "POSTGRES_PASSWORD=" + p.Password,
		"-e", "POSTGRES_DB=" + p.Database,
		"--publish=127.0.0.1:0:5432/tcp",
		"postgres:9.6-alpine",
	}
	cmd := exec.Command("docker", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	dry.Check(cmd.Run())

	cmd = exec.Command("docker", "container", "port", p.ContainerName, "5432/tcp")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	dry.Check(err)
	p.Port = strings.SplitN(strings.TrimSpace(string(out)), ":", 2)[1]

	return &p
}

func (p *PostgreSQL) DSN() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres", // user
		p.Password,
		"127.0.0.1", // host
		p.Port,
		p.Database,
	)
}

type tabler interface {
	TableName() string
}

func (p *PostgreSQL) Dump(allModels []interface{}) []byte {
	tables := make([]string, 0, 2*len(allModels))
	for _, m := range allModels {
		t := m.(tabler).TableName()
		tables = append(tables, "-t", t)
	}
	args := []string{
		"exec",
		"--user=postgres",
		p.ContainerName,
		"pg_dump", "-s", "-xO", p.Database,
	}
	args = append(args, tables...)
	log.Println(args)
	cmd := exec.Command("docker", args...)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	dry.Check(err)
	return out
}

func (p *PostgreSQL) Remove() {
	log.Println("stopping postgresql container")
	cmd := exec.Command("docker", "container", "rm", "-f", p.ContainerName)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	dry.Check(cmd.Run())
}
