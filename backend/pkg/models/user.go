package models

import (
	octopus "backend/app"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"html"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users []User

type User struct {
	ID          uuid.UUID    `sql:"type:uuid;primary key" json:"id"`
	Email       string       `sql:"type:varchar(100);unique" json:"email"`
	Password    string       `sql:"type:varchar(100)" json:"password"`
	FirstName   string       `sql:"type:varchar(100)" json:"firstName"`
	LastName    string       `sql:"type:varchar(100)" json:"lastName"`
	DateOfBirth time.Time    `json:"dateOfBirth"`
	AvatarImage string       `sql:"type:varchar(255)" json:"avatarImage"`
	Nickname    string       `sql:"type:varchar(100);unique" json:"nickname"`
	AboutMe     string       `sql:"type:text" json:"aboutMe"`
	IsPublic    bool         `json:"isPublic"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   sql.NullTime `json:"deletedAt"`
}

// Create a new user
func (user *User) Create(db *sql.DB) error {

	// Define the user default properties
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	paswordCrypted, err := bcrypt.GenerateFromPassword([]byte(html.EscapeString(user.Password)), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userFeildsCheckCases := []bool{
		user.Email != "",
		user.Password != "",
		user.Nickname != "",
		user.FirstName != "",
		user.LastName != "",
	}
	for _, v := range userFeildsCheckCases {
		if !v {
			return errors.New("some empty values in form")
		}
	}
	query := `INSERT INTO users (id, email, password, first_name, last_name, date_of_birth, avatar_image, nickname, about_me, is_public, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`

	stmt, err := db.Prepare(query)
	if err != nil {

		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		user.ID.String(),
		html.EscapeString(user.Email),
		paswordCrypted,
		html.EscapeString(user.FirstName),
		html.EscapeString(user.LastName),
		user.DateOfBirth,
		html.EscapeString(user.AvatarImage),
		html.EscapeString(user.Nickname),
		html.EscapeString(user.AboutMe),
		user.IsPublic,
		user.CreatedAt,
		user.UpdatedAt,
	)
	if err != nil {
		fmt.Println("here")
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Get a user by its ID
func (user *User) Get(db *sql.DB, id uuid.UUID) error {
	query := `SELECT id, email, password, first_name, last_name, date_of_birth, avatar_image, nickname, about_me, is_public, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL`

	row := db.QueryRow(query, id.String())

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

// Update a user
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

// Delete a user
func (user *User) Delete(db *sql.DB) error {
	query := `UPDATE users SET deleted_at=$1 WHERE id=$2`

	stmt, err := db.Prepare(query)
	if err != nil {
	}
	return fmt.Errorf("unable to prepare the query. %v", err)
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), user.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	return nil
}

func (user *User) CheckCredentials(ctx *octopus.Context) bool {
	query := `SELECT id,password FROM users WHERE (email = ? OR nickname = ?)`
	log.Println("✅>>>>>>>>>>>>>>>>>>>>", user.Email, user.Password)

	var realPassword string
	err := ctx.Db.Conn.QueryRow(query, user.Email, user.Email).Scan(&user.ID, &realPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(realPassword), []byte(user.Password))

	return err == nil
}

// GetAll users
func (users *Users) GetAll(db *sql.DB) error {
	query := `SELECT id, email, password, first_name, last_name, date_of_birth, avatar_image, nickname, about_me, is_public, created_at, updated_at FROM users WHERE deleted_at IS NULL`

	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(
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
			return fmt.Errorf("unable to execute the query. %v", err)
		}
		*users = append(*users, user)
	}

	return nil
}
