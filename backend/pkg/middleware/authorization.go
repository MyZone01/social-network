package middleware

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/models"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func AuthMiddleware(ctx *octopus.Context) {
	// AuthMiddleware is responsible for checking if the incoming request is authenticated.
	// It uses the 'ctx' object to access request data and perform authentication checks.
	// If the middleware determines that the request is not authenticated, it should return early.
	// Otherwise, it calls 'ctx.Next()' to pass control to the next middleware or handler in the chain.

	var middlewarePassed bool
	// TODO: Implement the actual authentication logic here.
	// For example, check for a valid session token or API key in the request headers.
	// If the check passes, set 'middlewarePassed' to true.

	if !middlewarePassed {
		ctx.WriteString("the middleware did not pass ")
		// If the middleware did not pass, return without calling 'ctx.Next()'.
		// This effectively stops the request from reaching subsequent middleware or handlers.
		return
	}

	// If the middleware passed, call 'ctx.Next()' to continue processing the request.
	ctx.Next()
}

// AuthAccessGoup is a middleware that checks if the user is authenticated
func AuthAccessGoup(ctx *octopus.Context) {
	var token string
	headerBearer := ctx.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	}

	user, err := config.Sess.Start(ctx).Get(token)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	var data = map[string]interface{}{}
	if err := ctx.BodyParser(&data); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Erreur de données.",
		})
		return
	}

	groupId, ok := data["group_id"].(string)
	if !ok {
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Erreur de données.",
		})
		return
	}

	var mg = new(models.GroupMember)
	if err := mg.GetMember(ctx.Db.Conn, user, uuid.MustParse(groupId), false); err != nil {
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

// CreatePostMiddleware is a middleware that checks if the data is valid
func CreatePostMiddleware(c *octopus.Context) {
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
		data["Title"] == nil || strings.TrimSpace(data["Title"].(string)) == "" ||
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
func CreateGroupPostMiddleware(c *octopus.Context) {
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
		data["content"] == nil || strings.TrimSpace(data["content"].(string)) == "" ||
		(data["image_url"] == nil && strings.TrimSpace(data["image_url"].(string)) == "") {
		c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid data",
		})
		return
	}
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