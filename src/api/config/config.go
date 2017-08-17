package config

import (
	"os"
)

const (
	DEVELOPMENT = "development"
	STAGING     = "staging"
	PRODUCTION  = "production"
)

type config struct {
	ReleaseStage string

	DatabaseDSN string
	QueryLog    bool

	BugSnagKey string
}

var Config config

func init() {
	Config.ReleaseStage = os.Getenv("RELEASE_STAGE")
	if Config.ReleaseStage == "" {
		Config.ReleaseStage = DEVELOPMENT
	}
	switch Config.ReleaseStage {
	case PRODUCTION:
		Config.DatabaseDSN = os.Getenv("DATABASE_DSN")
		Config.QueryLog = false

	case DEVELOPMENT, STAGING:
		Config.QueryLog = Config.ReleaseStage == DEVELOPMENT
	}

	Config.BugSnagKey = "..."
}
