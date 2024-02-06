package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PrivateMessage struct {
	ID         uuid.UUID `sql:"type:uuid;primary key"`
	SenderID   uuid.UUID `sql:"type:uuid"`
	ReceiverID uuid.UUID `sql:"type:uuid"`
	Content    string    `sql:"type:text"`
	CreatedAt  time.Time
	DeletedAt  sql.NullTime
}

type GroupMessage struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	SenderID  uuid.UUID `sql:"type:uuid"`
	Content   string    `sql:"type:text"`
	CreatedAt time.Time
	DeletedAt sql.NullTime
}
