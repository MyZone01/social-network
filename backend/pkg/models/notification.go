package models

import (
	"time"

	"github.com/google/uuid"
)

type NotificationType string

const (
	TypeFollowRequest   NotificationType = "follow_request"
	TypeGroupInvitation NotificationType = "group_invitation"
	TypeNewMessage      NotificationType = "new_message"
	TypeNewEvent        NotificationType = "new_event"
	// Add more types as needed
)

type Notification struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	UserID    uuid.UUID `sql:"type:uuid"`
	Type      NotificationType
	Message   string `sql:"type:text"`
	CreatedAt time.Time
}
