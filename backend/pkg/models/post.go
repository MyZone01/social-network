package models

import (
	"database/sql"
	"fmt"
	"html"
	"time"

	"github.com/google/uuid"
)

type PostPrivacy string
type Posts []Post

const (
	PrivacyPublic        PostPrivacy = "public"
	PrivacyPrivate       PostPrivacy = "private"
	PrivacyAlmostPrivate PostPrivacy = "almost private"
	PrivacyUnlisted      PostPrivacy = "unlisted"
)

type Post struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"userId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	ImageURL  string    `json:"image_url"`
	Privacy   PostPrivacy `json:"privacy"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

// IsPublic returns true if the post is public
func (p *Post) IsPublic() bool {
	return p.Privacy == PrivacyPublic
}

// IsPrivate returns true if the post is private
func (p *Post) IsPrivate() bool {
	return p.Privacy == PrivacyPrivate
}

// IsAlmostPrivate returns true if the post is almost private
func (p *Post) IsAlmostPrivate() bool {
	return p.Privacy == PrivacyAlmostPrivate
}

// IsUnlisted returns true if the post is unlisted
func (p *Post) IsUnlisted() bool {
	return p.Privacy == PrivacyUnlisted
}

// IsDeleted returns true if the post is deleted
func (p *Post) IsDeleted() bool {
	return p.DeletedAt.Valid
}

// PostPrivacyFromString returns the post privacy from a string
func PostPrivacyFromString(s string) (PostPrivacy, error) {
	switch s {
	case "public":
		return PrivacyPublic, nil
	case "private":
		return PrivacyPrivate, nil
	case "almost private":
		return PrivacyAlmostPrivate, nil
	case "unlisted":
		return PrivacyUnlisted, nil
	default:
		return "", fmt.Errorf("invalid post privacy")
	}
}

// Create inserts a new post into the database
func (p *Post) Create(db *sql.DB, userOwnerUuid uuid.UUID) error {
	// Define the post default properties
	p.ID = uuid.New()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.UserID = userOwnerUuid
	query := `INSERT INTO posts (id, user_id,title, content, image_url, privacy, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		p.ID,
		p.UserID.String(),
		html.EscapeString(p.Title),
		html.EscapeString(p.Content),
		html.EscapeString(p.ImageURL),
		p.Privacy,
		p.CreatedAt,
		p.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Get retrieves a post from the database
func (p *Post) Get(db *sql.DB, id uuid.UUID) error {
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE id = $1 AND deleted_at IS NULL`

	err := db.QueryRow(query, id).Scan(
		&p.ID,
		&p.UserID,
		&p.Title,
		&p.Content,
		&p.ImageURL,
		&p.Privacy,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no post found with id %v", id)
		}
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Update modifies a post in the database
func (p *Post) Update(db *sql.DB) error {
	query := `UPDATE posts SET title = $1, content = $2, image_url = $3, privacy = $4, updated_at = $5 WHERE id = $6`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		html.EscapeString(p.Title),
		html.EscapeString(p.Content),
		html.EscapeString(p.ImageURL),
		p.Privacy,
		time.Now(),
		p.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Delete removes a post from the database
func (p *Post) Delete(db *sql.DB) error {
	query := `UPDATE posts SET deleted_at = $1 WHERE id = $2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		time.Now(),
		p.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetUserPosts retrieves all the posts from a user
func (p *Posts) GetUserPosts(db *sql.DB, userID uuid.UUID) error {
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE user_id = $1 AND deleted_at IS NULL`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(userID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err := rows.Scan(
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
			return fmt.Errorf("unable to scan the row. %v", err)
		}
		*p = append(*p, post)
	}

	return nil
}

// GetAll retrieves all the posts from the database
func (p *Posts) GetAll(db *sql.DB) error {
	query := `SELECT id, user_id, title, content, image_url, privacy, created_at, updated_at, deleted_at FROM posts WHERE deleted_at IS NULL`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	for rows.Next() {
		var post Post
		err := rows.Scan(
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
			return fmt.Errorf("unable to scan the row. %v", err)
		}
		*p = append(*p, post)
	}

	return nil
}
