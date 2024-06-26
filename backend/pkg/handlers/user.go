package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type updateValues struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"newpassword"`
}

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

func handleUpdateUserInfos(ctx *octopus.Context) {
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

	idSession, err := config.Sess.Start(ctx).Set(user.ID)
	// Start a new session for the user and set the user's ID as the session key.
	if err != nil {
		// If starting the session fails, log the error and return an HTTP status  500.
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while starting the session.",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	data := map[string]interface{}{
		"message": "User updated successfully",
		"session": idSession,
		"user":    user,
		"status":  http.StatusOK,
	}
	ctx.JSON(data)
}

var updateUserInfosRoute = route{
	path:   "/edituser",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleUpdateUserInfos,   // Handler function to process the authentication request.
	},
}

func handleUpdateUserPassword(ctx *octopus.Context) {
	// Log the client's IP address that reached the login route.
	var newCredentials = updateValues{}

	// Try to deserialize the form data into the User instance.
	if err := ctx.BodyParser(&newCredentials); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while parsing the form data.",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	var credentials = credentials{
		Email:    newCredentials.Email,
		Password: newCredentials.Password,
	}

	if err := credentials.Validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	newUser := models.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	err := newUser.Get(ctx.Db.Conn, credentials.Email)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "invalid email.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	// Check if the user's credentials are valid.
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(credentials.Password)); err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "Invalid credentials. Please try again.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(newCredentials.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while hashing the password.",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	newUser.Password = string(newPasswordHash)

	// Attempts to update a the user in the database with the provided data.
	if newUser.Update(ctx.Db.Conn) != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while creating the user.",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	// Starts a new session for the user and sets the user's ID as the session key.
	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	if err != nil {
		// If starting the session fails, logs the error and returns an HTTP status  500.
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while starting the session.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.Status(http.StatusAccepted).JSON(map[string]interface{}{
		"session": idSession,
		"user":    newUser,
		"message": "User successfully registered and logged.",
		"status":  "200",
	})
}

var updateUserPasswordRoute = route{
	path:   "/updatepassword",
	method: http.MethodPut,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,  // Middleware to check if the request is authenticated.
		handleUpdateUserPassword, // Handler function to process the authentication request.
	},
}

func init() {
	AllHandler[updateUserRoute.path] = updateUserRoute
	AllHandler[updateUserInfosRoute.path] = updateUserInfosRoute
	AllHandler[updateUserPasswordRoute.path] = updateUserPasswordRoute
}
