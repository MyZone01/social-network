package models

import (
	"database/sql"
	"fmt"
	"html"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `sql:"type:uuid;primary key"`
	Email       string    `sql:"type:varchar(100);unique"`
	Password    string    `sql:"type:varchar(100)"`
	FirstName   string    `sql:"type:varchar(100)"`
	LastName    string    `sql:"type:varchar(100)"`
	DateOfBirth time.Time
	AvatarImage string `sql:"type:varchar(255)"`
	Nickname    string `sql:"type:varchar(100);unique"`
	AboutMe     string `sql:"type:text"`
	IsPublic    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

func (user *User) Create(db *sql.DB) error {
	query := `INSERT INTO users (id, email, password, first_name, last_name, date_of_birth, avatar_image, nickname, about_me, is_public, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.ID,
		html.EscapeString(user.Email),
		html.EscapeString(user.Password),
		html.EscapeString(user.FirstName),
		html.EscapeString(user.LastName),
		user.DateOfBirth,
		html.EscapeString(user.AvatarImage),
		html.EscapeString(user.Nickname),
		html.EscapeString(user.AboutMe),
		user.IsPublic,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

func (user *User) Get(db *sql.DB, id uuid.UUID) error {
	query := `SELECT id, email, password, first_name, last_name, date_of_birth, avatar_image, nickname, about_me, is_public, created_at, updated_at FROM users WHERE id = $1`

	row := db.QueryRow(query, id)

	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.AvatarImage,
		&user.Nickname,
		&user.AboutMe,
		&user.IsPublic,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no user found with id %v", id)
		}
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

func (user *User) Update(db *sql.DB) error {
	query := `UPDATE users SET email=$1, password=$2, first_name=$3, last_name=$4, date_of_birth=$5, avatar_image=$6, nickname=$7, about_me=$8, is_public=$9, updated_at=$10 WHERE id=$11`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		html.EscapeString(user.Email),
		html.EscapeString(user.Password),
		html.EscapeString(user.FirstName),
		html.EscapeString(user.LastName),
		user.DateOfBirth,
		html.EscapeString(user.AvatarImage),
		html.EscapeString(user.Nickname),
		html.EscapeString(user.AboutMe),
		user.IsPublic,
		time.Now(),
		user.ID,
	)

	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

func (user *User) Delete(db *sql.DB) error {
	query := `DELETE FROM users WHERE id=$1`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}
