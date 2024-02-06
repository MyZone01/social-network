package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `sql:"type:uuid;primary key"`
	GroupID     uuid.UUID `sql:"type:uuid"`
	Title       string    `sql:"type:varchar(255)"`
	Description string    `sql:"type:text"`
	DateTime    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type EventResponse string

const (
	ResponseGoing    EventResponse = "going"
	ResponseNotGoing EventResponse = "not_going"
)

type EventParticipant struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	MemberID  uuid.UUID `sql:"type:uuid"`
	Response  EventResponse
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
