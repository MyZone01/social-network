package models

import (
	"database/sql"
	"fmt"
	"html"
	"time"

	"github.com/google/uuid"
)

type Groups []Group
type GroupMembers []GroupMember
type GroupPosts []GroupPost

type Group struct {
	ID           uuid.UUID `sql:"type:uuid;primary key"`
	Title        string    `sql:"type:varchar(255)"`
	Description  string    `sql:"type:text"`
	BannerURL    string    `sql:"type:varchar(255)"`
	CreatorID    uuid.UUID `sql:"type:uuid"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime
	GroupMembers GroupMembers
}

type GroupMemberStatus string

const (
	MemberStatusInvited    GroupMemberStatus = "invited"
	MemberStatusRequesting GroupMemberStatus = "requesting"
	MemberStatusAccepted   GroupMemberStatus = "accepted"
	MemberStatusDeclined   GroupMemberStatus = "declined"
)

type GroupMemberRole string

const (
	MemberRoleAdmin GroupMemberRole = "admin"
	MemberRoleUser  GroupMemberRole = "user"
)

type GroupMember struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	MemberID  uuid.UUID `sql:"type:uuid"`
	Status    GroupMemberStatus
	Role      GroupMemberRole
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	User      User
}

type GroupPost struct {
	ID        uuid.UUID `sql:"type:uuid;primary key"`
	CreatorID uuid.UUID `sql:"type:uuid"`
	GroupID   uuid.UUID `sql:"type:uuid"`
	PostID    uuid.UUID `sql:"type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Post      Post
}

// Create inserts a new group into the database
func (g *Group) Create(db *sql.DB) error {
	// Define the group default properties
	g.ID = uuid.New()
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	query := `INSERT INTO groups (id, title, description, banner_url, creator_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(g.ID, html.EscapeString(g.Title), html.EscapeString(g.Description), html.EscapeString(g.BannerURL), g.CreatorID, g.CreatedAt, g.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	// Create the group creator as a admin of the group
	gm := GroupMember{
		Status: MemberStatusAccepted,
		Role:   MemberRoleAdmin,
	}
	err = gm.CreateMember(db, g.CreatorID, g.ID)
	if err != nil {
		return fmt.Errorf("unable to create group member. %v", err)
	}

	g.GroupMembers = append(g.GroupMembers, gm)

	return nil
}

// Get retrieves a group from the database
func (g *Group) Get(db *sql.DB, id uuid.UUID, getmembers, getuser bool) error {
	query := `SELECT id, title, description, banner_url, creator_id, created_at, updated_at, deleted_at FROM groups WHERE id=$1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	row := stm.QueryRow(id)
	err = row.Scan(
		&g.ID,
		&g.Title,
		&g.Description,
		&g.BannerURL,
		&g.CreatorID,
		&g.CreatedAt,
		&g.UpdatedAt,
		&g.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("unable to scan the row. %v", err)
	}

	if getmembers {
		err = g.GetMembers(db, getuser)
		if err != nil {
			return fmt.Errorf("unable to get group members. %v", err)
		}
	}

	return nil
}

// Update updates the group in the database
func (g *Group) Update(db *sql.DB) error {
	g.UpdatedAt = time.Now()
	query := `UPDATE groups SET title=$1, description=$2, banner_url=$3, updated_at=$4 WHERE id=$5`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(html.EscapeString(g.Title), html.EscapeString(g.Description), html.EscapeString(g.BannerURL), g.UpdatedAt, g.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// Delete removes the group from the database
func (g *Group) Delete(db *sql.DB) error {
	query := `UPDATE groups SET deleted_at=$1 WHERE id=$2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), g.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetAllGroups retrieves all groups from the database
func (gs *Groups) GetAllGroups(db *sql.DB, getmembers, getuser bool) error {
	query := `SELECT id, title, description, banner_url, creator_id, created_at, updated_at, deleted_at FROM groups WHERE deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query()
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var g Group
		err = rows.Scan(
			&g.ID,
			&g.Title,
			&g.Description,
			&g.BannerURL,
			&g.CreatorID,
			&g.CreatedAt,
			&g.UpdatedAt,
			&g.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		if getmembers {
			err = g.GetMembers(db, getuser)
			if err != nil {
				return fmt.Errorf("unable to get group members. %v", err)
			}
		}

		*gs = append(*gs, g)
	}

	return nil
}

// CreateMember inserts a new member into the group in the database
func (gm *GroupMember) CreateMember(db *sql.DB, memberID, groupID uuid.UUID) error {
	gm.ID = uuid.New()
	gm.GroupID = groupID
	gm.MemberID = memberID
	gm.CreatedAt = time.Now()
	gm.UpdatedAt = time.Now()
	query := `INSERT INTO groupMembers (id, group_id, member_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	_, err = stmt.Exec(gm.ID, gm.GroupID, gm.MemberID, gm.Status, gm.CreatedAt, gm.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetMember retrieves a member from the group in the database
func (gm *GroupMember) GetMember(db *sql.DB, memberID, groupID uuid.UUID, getuser bool) error {
	query := `SELECT id, group_id, member_id, status, created_at, updated_at, deleted_at FROM groupMembers WHERE group_id=$1 AND member_id=$2 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	row := stm.QueryRow(groupID, memberID)
	err = row.Scan(
		&gm.ID,
		&gm.GroupID,
		&gm.MemberID,
		&gm.Status,
		&gm.CreatedAt,
		&gm.UpdatedAt,
		&gm.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("unable to scan the row. %v", err)
	}

	if getuser {
		var user = new(User)
		err = user.Get(db, memberID)
		if err != nil {
			return fmt.Errorf("unable to get user. %v", err)
		}
	}

	return nil
}

// UpdateMember updates the member in the group in the database
func (gm *GroupMember) UpdateMember(db *sql.DB) error {
	gm.UpdatedAt = time.Now()
	query := `UPDATE groupMembers SET status=$1, updated_at=$2 WHERE id=$3`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(gm.Status, gm.UpdatedAt, gm.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// DeleteMember removes the member from the group in the database
func (gm *GroupMember) DeleteMember(db *sql.DB) error {
	query := `UPDATE groupMembers SET deleted_at=$1 WHERE id=$2`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), gm.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetMembers retrieves all members of the group from the database
func (g *Group) GetMembers(db *sql.DB, getusers bool) error {
	query := `SELECT id, group_id, member_id, status, created_at, updated_at, deleted_at FROM groupMembers WHERE group_id=$1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(g.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var gm GroupMember
		err = rows.Scan(
			&gm.ID,
			&gm.GroupID,
			&gm.MemberID,
			&gm.Status,
			&gm.CreatedAt,
			&gm.UpdatedAt,
			&gm.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		if getusers {
			var user = new(User)
			err = user.Get(db, gm.MemberID)
			if err != nil {
				return fmt.Errorf("unable to get user. %v", err)
			}
			gm.User = *user
		}

		g.GroupMembers = append(g.GroupMembers, gm)
	}

	return nil
}

// CreatePost inserts a new post into the group in the database
func (gp *GroupPost) CreatePost(db *sql.DB,) error {
	// Define the group post default properties
	gp.ID = uuid.New()
	gp.CreatedAt = time.Now()
	gp.UpdatedAt = time.Now()
	gp.Post.CreatedAt = time.Now()
	gp.Post.UpdatedAt = time.Now()
	gp.Post.ID = uuid.New()
	gp.Post.UserID = gp.ID
	gp.PostID = gp.Post.ID

	if err := gp.Post.Create(db, gp.CreatorID); err != nil {
		return fmt.Errorf("unable to create the post. %v", err)
	}

	query := `INSERT INTO groupPosts (id, group_id, post_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(gp.ID, gp.GroupID, gp.PostID, gp.CreatedAt, gp.UpdatedAt)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetPost retrieves a post from the group in the database
func (gp *GroupPost) GetPost(db *sql.DB, groupID, groupPostID uuid.UUID, getpost bool) error {
	query := `SELECT id, group_id, post_id, created_at, updated_at, deleted_at FROM groupPosts WHERE group_id=$1 AND post_id=$2 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	row := stm.QueryRow(groupID, groupPostID)
	err = row.Scan(
		&gp.ID,
		&gp.GroupID,
		&gp.PostID,
		&gp.CreatedAt,
		&gp.UpdatedAt,
		&gp.DeletedAt,
	)
	if err != nil {
		return fmt.Errorf("unable to scan the row. %v", err)
	}

	if getpost {
		err = gp.Post.Get(db, gp.PostID)
		if err != nil {
			return fmt.Errorf("unable to get post. %v", err)
		}
	}

	return nil
}

// UpdatePost updates the post in the group in the database
func (gp *GroupPost) UpdatePost(db *sql.DB) error {
	gp.UpdatedAt = time.Now()
	query := `UPDATE groupPosts SET updated_at=$1 WHERE id=$2`

	// Update the post
	if err := gp.Post.Update(db); err != nil {
		return fmt.Errorf("unable to update the post. %v", err)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(gp.UpdatedAt, gp.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// DeletePost removes the post from the group in the database
func (gp *GroupPost) DeletePost(db *sql.DB) error {
	query := `UPDATE groupPosts SET deleted_at=$1 WHERE id=$2`

	// Delete the post
	if err := gp.Post.Delete(db); err != nil {
		return fmt.Errorf("unable to delete the post. %v", err)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(time.Now(), gp.ID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}

	return nil
}

// GetPosts retrieves all posts of the group from the database
func (g *GroupPosts) GetPosts(db *sql.DB, groupID uuid.UUID, getpost bool) error {
	query := `SELECT id, group_id, post_id, created_at, updated_at, deleted_at FROM groupPosts WHERE group_id=$1 AND deleted_at IS NULL`

	stm, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("unable to prepare the query. %v", err)
	}
	defer stm.Close()

	rows, err := stm.Query(groupID)
	if err != nil {
		return fmt.Errorf("unable to execute the query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var gp GroupPost
		err = rows.Scan(
			&gp.ID,
			&gp.GroupID,
			&gp.PostID,
			&gp.CreatedAt,
			&gp.UpdatedAt,
			&gp.DeletedAt,
		)
		if err != nil {
			return fmt.Errorf("unable to scan the row. %v", err)
		}

		if getpost {
			err = gp.GetPost(db, gp.GroupID, gp.PostID, true)
			if err != nil {
				return fmt.Errorf("unable to get post. %v", err)
			}
		}

		*g = append(*g, gp)
	}

	return nil
}
