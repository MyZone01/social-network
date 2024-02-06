package models

import (
	"time"

	"github.com/google/uuid"
)

type FollowerStatus string

const (
	StatusRequested FollowerStatus = "requested"
	StatusAccepted  FollowerStatus = "accepted"
	StatusDeclined  FollowerStatus = "declined"
)

type Follower struct {
	ID         uuid.UUID `sql:"type:uuid;primary key"`
	FollowerID uuid.UUID `sql:"type:uuid"`
	FolloweeID uuid.UUID `sql:"type:uuid"`
	Status     FollowerStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
