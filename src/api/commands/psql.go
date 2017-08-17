package commands

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/urfave/cli"

	"api/config"
	"api/dry"
)

//noinspection GoUnusedParameter
func Psql(c *cli.Context) error {
	u, err := url.Parse(config.Config.DatabaseDSN)
	dry.Check(err)

	host := strings.SplitN(u.Host, ":", 2)
	password, _ := u.User.Password()
	env := []string{
		"PGHOST=" + host[0],
		"PGUSER=" + u.User.Username(),
		"PGPASSWORD=" + password,
		"PGPORT=" + host[1],
		"PGDATABASE=" + u.Path[1:],
	}

	fmt.Println(strings.Join(env, " "), "psql")
	return nil
}
