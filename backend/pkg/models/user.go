package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID      `sql:"type:uuid;primary key"`
	Email        string         `sql:"type:varchar(100);unique"`
	Password     string         `sql:"type:varchar(100)"`
	FirstName    string         `sql:"type:varchar(100)"`
	LastName     string         `sql:"type:varchar(100)"`
	DateOfBirth  time.Time
	AvatarImage  string         `sql:"type:varchar(255)"`
	Nickname     string         `sql:"type:varchar(100);unique"`
	AboutMe      string         `sql:"type:text"`
	IsPublic     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
}