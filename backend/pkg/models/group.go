package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Group struct {
	ID          uuid.UUID `sql:"type:uuid;primary key"`
	Title       string    `sql:"type:varchar(255)"`
	Description string    `sql:"type:text"`
	BannerURL   string    `sql:"type:varchar(255)"`
	CreatorID   uuid.UUID `sql:"type:uuid"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

type GroupMemberStatus string

const (
	MemberStatusInvited    GroupMemberStatus = "invited"
	MemberStatusRequesting GroupMemberStatus = "requesting"
	MemberStatusAccepted   GroupMemberStatus = "accepted"
	MemberStatusDeclined   GroupMemberStatus = "declined"
)

type GroupMember struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	MemberID  uuid.UUID `sql:"type:uuid"`
	Status    GroupMemberStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type GroupPost struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	PostID    uuid.UUID `sql:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
