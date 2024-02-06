package models

import (
	"database/sql"
	"html"
	"time"

	"github.com/google/uuid"
)

type NotificationType string
type Notifications []Notification

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
	DeletedAt sql.NullTime
}

// Create a new notification
func (n *Notification) Create(db *sql.DB) error {
	n.ID = uuid.New()
	n.CreatedAt = time.Now()

	query := `INSERT INTO notifications (id, user_id, type, message, created_at) 
		VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(query, n.ID, n.UserID, n.Type, html.EscapeString(n.Message), n.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

// Get a notification by its ID
func (n *Notification) Get(db *sql.DB, id uuid.UUID) error {
	query := `SELECT id, user_id, type, message, created_at, deleted_at FROM notifications WHERE id = $1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	err = stm.QueryRow(id).Scan(&n.ID, &n.UserID, &n.Type, &n.Message, &n.CreatedAt, &n.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

// Delete a notification
func (n *Notification) Delete(db *sql.DB) error {
	query := `UPDATE notifications SET deleted_at = $1 WHERE id = $2`

	_, err := db.Exec(query, time.Now(), n.ID)
	if err != nil {
		return err
	}
	return nil
}

// Get all notifications for a user
func (n *Notifications) GetByUser(db *sql.DB, userID uuid.UUID) error {
	query := `SELECT id, user_id, type, message, created_at, deleted_at FROM notifications WHERE user_id = $1 AND deleted_at IS NULL`

	stm , err := db.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()

	rows, err := stm.Query(userID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var notif Notification
		err = rows.Scan(&notif.ID, &notif.UserID, &notif.Type, &notif.Message, &notif.CreatedAt, &notif.DeletedAt)
		if err != nil {
			return err
		}
		*n = append(*n, notif)
	}

	return nil
}

// ClearByUser deletes all notifications for a user
func (n *Notifications) ClearByUser(db *sql.DB, userID uuid.UUID) error {
	query := `UPDATE notifications SET deleted_at = $1 WHERE user_id = $2`

	_, err := db.Exec(query, time.Now(), userID)
	if err != nil {
		return err
	}
	return nil
}
