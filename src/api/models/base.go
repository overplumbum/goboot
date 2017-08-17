package models

import (
	"time"

	"api/config"
	"api/dry"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satori/go.uuid"
)

type UUIDModel struct {
	ID uuid.UUID `sql:"TYPE:uuid;PRIMARY_KEY;NOT NULL;DEFAULT:uuid_generate_v4()"`
}

func (e *UUIDModel) ApiId() *string {
	if e == nil || e.ID == uuid.Nil {
		return nil
	} else {
		out := e.ID.String()
		return &out
	}
}

func (e *UUIDModel) Empty() bool {
	return e == nil || e.ID == uuid.Nil
}

type BaseModel struct {
	UUIDModel

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

var DB *gorm.DB

func SetupDB() {
	var err error
	DB, err = gorm.Open("postgres", config.Config.DatabaseDSN)
	dry.Check(err)

	DB.SingularTable(true)
	if config.Config.QueryLog {
		DB.LogMode(true)
	}
}

// full models list to be created during migration
var Models []interface{}

func Migrate() {
	dry.Check(DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error)
	dry.Check(DB.AutoMigrate(Models...).Error)
}
