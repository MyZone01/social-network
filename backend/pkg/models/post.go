package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
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

func (post *Post) Create(db *sql.DB) error {
	query := `INSERT INTO posts (id, user_id, title, content, image_url, privacy, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		post.ID,
		post.UserID,
		html.EscapeString(post.Title),
		html.EscapeString(post.Content),
		html.EscapeString(post.ImageURL),
		post.Privacy,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

func (post *Post) Get(db *sql.DB, id uuid.UUID) error {
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE id = $1`

	row := db.QueryRow(query, id)

	err := row.Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Content,
		&post.ImageURL,
		&post.Privacy,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no post found with id %v", id)
		}
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

func (post *Post) Update(db *sql.DB) error {
	query := `UPDATE posts SET user_id=$1, title=$2, content=$3, image_url=$4, privacy=$5, updated_at=$6 WHERE id=$7`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		post.UserID,
		post.Title,
		post.Content,
		post.ImageURL,
		post.Privacy,
		time.Now(),
		post.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

func (post *Post) Delete(db *sql.DB) error {
	query := `DELETE FROM posts WHERE id=$1`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	return nil
}

func (post *Post) UnmarshalFormData(formData []byte) error {
	if err := json.Unmarshal(formData, post); err != nil {
		return err
	}
	return nil
}
