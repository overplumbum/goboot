package models

import (
	_ "github.com/jinzhu/gorm"
)

func init() {
	Models = append(
		Models,
		&Session{},
	)
}

type Session struct {
	BaseModel
}

func (Session) TableName() string {
	return "session"
}
