package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"
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

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var updateUserRoute = route{
	path:   "/updateuser",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		func(ctx *octopus.Context) {
			fmt.Println("fgr")
			ctx.Next()
		},
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleUpdateUser,         // Handler function to process the authentication request.
	},
}

func init() {
	AllHandler[updateUserRoute.path] = updateUserRoute
}
