package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `sql:"type:uuid;primary key"`
	UserID       uuid.UUID `sql:"type:uuid"`
	SessionToken string    `sql:"type:varchar(255)"`
	CreatedAt    time.Time
	ExpiresAt    time.Time
	DeletedAt    sql.NullTime
}
