package middleware

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/models"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// HaveGroupAccess is a middleware that checks if the user is authenticated
func HaveGroupAccess(ctx *octopus.Context) {
	groupId := ctx.Request.URL.Query().Get("group_id")
	if groupId == "" {
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Group Id",
		})
		return
	}

	userUUID := ctx.Values["userId"].(uuid.UUID)

	var mg = new(models.GroupMember)
	if err := mg.GetMember(ctx.Db.Conn, userUUID, uuid.MustParse(groupId), false); err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas autorisé.",
		})
		return
	}
	ctx.Values["role"] = mg.Role
	ctx.Next()
}

// CheckRole is a middleware that checks if the user have a specific role in the group
func CheckGroupRole(ctx *octopus.Context, role models.GroupMemberRole) {
	_role, ok := ctx.Values["role"].(models.GroupMemberRole)
	if !ok {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas autorisé.",
		})
		return
	}
	if _role != role {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas autorisé.",
		})
		return
	}
	ctx.Next()
}

// ImageUploadMiddleware is a middleware that checks if the file is an image and downloads it
func ImageUploadMiddleware(c *octopus.Context) {
	c.Request.ParseMultipartForm(10 << 20) // limit your max input length!
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Error Retrieving the File",
		})
		return
	}
	defer file.Close()
	// Check if the file is an image
	ext := []string{".jpeg", ".jpg", ".png", ".svg+xml"}

	if !contains(ext, strings.ToLower(filepath.Ext(handler.Filename))) {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "File type not allowed",
		})
		return
	}

	id := uuid.New()
	path := path.Join("uploads", id.String()+filepath.Ext(handler.Filename))
	// Create the file using the id as the name and the extension from the original file
	dst, err := os.Create(path)
	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "Error creating the file",
		})
		return
	}

	// Write the file
	if _, err := io.Copy(dst, file); err != nil {
		c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "Error writing file",
		})
		return
	}
	c.Values["file"] = path
	c.Next()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// IsPostValid is a middleware that checks if the data is valid
func IsPostValid(c *octopus.Context) {
	var data = map[string]interface{}{}

	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}
	privacy := data["privacy"].(string)
	_, err := models.PostPrivacyFromString(privacy)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid privacy",
		})
		return
	}

	if data["content"] == nil || strings.TrimSpace(data["content"].(string)) == "" ||
		data["title"] == nil || strings.TrimSpace(data["title"].(string)) == "" ||
		(data["image_url"] == nil && strings.TrimSpace(data["image_url"].(string)) == "") {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid data",
		})
		return
	}
	c.Next()
}

// CreateGroupMiddleware is a middleware that checks if the data is valid
func CreateGroupMiddleware(c *octopus.Context) {
	var token string
	headerBearer := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	var data = map[string]interface{}{}

	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}
	_, err := config.Sess.Start(c).Get(token)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	if data["title"] == nil || strings.TrimSpace(data["title"].(string)) == "" ||
		data["description"] == nil || strings.TrimSpace(data["description"].(string)) == "" ||
		(data["banner_url"] == nil && strings.TrimSpace(data["banner_url"].(string)) == "") {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid data",
		})
		return
	}
	c.Next()
}

// CreateEventMiddleware is a middleware that checks if the data is valid
func CreateEventMiddleware(c *octopus.Context) {
	var token string
	headerBearer := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	var data = map[string]interface{}{}

	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}
	_, err := config.Sess.Start(c).Get(token)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	if data["group_id"] == nil || strings.TrimSpace(data["group_id"].(string)) == "" ||
		data["title"] == nil || strings.TrimSpace(data["title"].(string)) == "" ||
		data["description"] == nil || strings.TrimSpace(data["description"].(string)) == "" ||
		data["date_time"] == nil || strings.TrimSpace(data["date_time"].(string)) == "" {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid data",
		})
		return
	}
	c.Next()
}

// CreateGroupPostMiddleware is a middleware that checks if the data is valid
func IsGroupPostValid(c *octopus.Context) {
	var token string
	headerBearer := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	var data = map[string]interface{}{}

	if err := c.BodyParser(&data); err != nil {
		log.Println(err.Error())
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}
	_, err := config.Sess.Start(c).Get(token)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	if data["title"] == nil || strings.TrimSpace(data["title"].(string)) == "" ||
		data["content"] == nil || strings.TrimSpace(data["content"].(string)) == "" || data["privacy"] == nil || strings.TrimSpace(data["privacy"].(string)) == "" || (data["privacy"] != "public" && data["privacy"] != "private") {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid data",
		})
		return
	}

	c.Next()
}

func IsGroupExist(c *octopus.Context) {
	_groupId := c.Request.URL.Query().Get("group_id")
	group := new(models.Group)
	// Check if the group is uuid

	groupId, err := uuid.Parse(_groupId)
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid group uuid",
		})
		return
	}

	if err := group.Get(c.Db.Conn, groupId, false, false); err != nil {
		c.Status(http.StatusNotFound).JSON(map[string]string{
			"error": "Group not found",
		})
		return
	}

	c.Values["group_id"] = groupId
	c.Next()
}

// CreateGroupMessageMiddleware is a middleware that checks if the data is valid
func CreateGroupMessageMiddleware(c *octopus.Context) {
	var token string
	headerBearer := c.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	var data = map[string]interface{}{}

	if err := c.BodyParser(&data); err != nil {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Error parsing request body",
		})
		return
	}
	_, err := config.Sess.Start(c).Get(token)
	if err != nil {
		c.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	if data["group_id"] == nil || strings.TrimSpace(data["group_id"].(string)) == "" ||
		data["content"] == nil || strings.TrimSpace(data["content"].(string)) == "" {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid data",
		})
		return
	}
	c.Next()
}
