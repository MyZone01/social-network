package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func handleUpdateUser(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)
	user := new(models.User)
	if err := ctx.BodyParser(user); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}
	user.ID = userId
	if err := user.Validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}
	newHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while hashing the password.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	user.Password = string(newHash)
	if err := user.Update(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.Status(http.StatusOK).JSON(map[string]interface{}{
		"message": "User updated successfully",
		"status":  http.StatusOK,
	})
}

func handleGetUser(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	type request struct {
		Action   string `json:"action"`
		Nickname string `json:"nickname"`
	}
	req := new(request)
	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}
	switch req.Action {
	case "get":
		user := new(models.User)
		if req.Nickname == "" {
			if err := user.Get(ctx.Db.Conn, userId); err != nil {
				ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
					"message": err.Error(),
					"status":  http.StatusInternalServerError,
				})
				return
			}
		} else {
			if err := user.Get(ctx.Db.Conn, req.Nickname); err != nil {
				ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
					"message": err.Error(),
					"status":  http.StatusInternalServerError,
				})
				return
			}
		}

		follower := new(models.Follower)
		follower.FollowerID = userId
		follower.FolloweeID = user.ID
		follower.Get(ctx.Db.Conn)

		if follower.Status == "" {
			follower.Status = "none"
		}
		if userId == user.ID {
			follower.Status = "self"
		}

		follow := new(models.Followers).CountAllByFollowerID(ctx.Db.Conn, userId)
		following := new(models.Followers).CountAllByFolloweeID(ctx.Db.Conn, userId)
		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "User fetched successfully",
			"status":  http.StatusOK,
			"data": map[string]interface{}{
				"id":           user.ID,
				"firstname":    user.FirstName,
				"lastname":     user.LastName,
				"email":        user.Email,
				"nickname":     user.Nickname,
				"birthday":     user.DateOfBirth,
				"about":        user.AboutMe,
				"avatar":       user.AvatarImage,
				"created":      user.CreatedAt,
				"updated":      user.UpdatedAt,
				"follow":       follow,
				"following":    following,
				"followStatus": follower.Status,
			},
		})
	case "posts":
		posts := new(models.Posts)
		if err := posts.GetUserPosts(ctx.Db.Conn, userId); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "User posts fetched successfully",
			"status":  http.StatusOK,
			"posts":   posts,
		})
	default:
		ctx.JSON(map[string]interface{}{
			"message": "Invalid action.",
			"status":  http.StatusBadRequest,
		})
	}
}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var updateUserRoute = route{
	path:   "/updateuser",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleUpdateUser,        // Handler function to process the authentication request.
	},
}

var getUserRoute = route{
	path:   "/getuser",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleGetUser,           // Handler function to process the authentication request.
	},
}

func init() {
	AllHandler[updateUserRoute.path] = updateUserRoute
	AllHandler[getUserRoute.path] = getUserRoute
}
