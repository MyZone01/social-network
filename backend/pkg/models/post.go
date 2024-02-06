package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PostPrivacy string

const (
	PrivacyPublic        PostPrivacy = "public"
	PrivacyPrivate       PostPrivacy = "private"
	PrivacyAlmostPrivate PostPrivacy = "almost private"
	PrivacyUnlisted      PostPrivacy = "unlisted"
)

type Post struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	UserID    uuid.UUID `sql:"type:uuid"`
	Title     string    `sql:"type:varchar(255)"`
	Content   string    `sql:"type:text"`
	ImageURL  string    `sql:"type:varchar(255)"`
	Privacy   PostPrivacy
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}
