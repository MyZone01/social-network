package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *credentials) Validate() error {
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New("Invalid email")
	}

	if c.Password == "" {
		return errors.New("Password is missing. Please provide it.")
	}

	return nil
}

// loginHandler is a function that handles user login requests.
// It attempts to unmarshal the form data from the client into a User instance,
// checks if the credentials are valid, and if successful, starts a new session for the user.
var loginHandler = func(ctx *octopus.Context) {
	// Log the client's IP address that reached the login route.
	var credentials = credentials{}

	// Try to deserialize the form data into the User instance.
	if err := ctx.BodyParser(&credentials); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while parsing the form data.",
			"status":  http.StatusInternalServerError,
		})
		return
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

	fmt.Println(newUser.Password, bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(credentials.Password)))

	// Check if the user's credentials are valid.
	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(credentials.Password)); err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "Invalid credentials. Please try again.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
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
	ctx.JSON(map[string]interface{}{
		"session": idSession,
		"message": "User successfully logged.",
		"status":  "200",
	})
}

// loginRoute is a structure that defines the login route for the API.
// It specifies that the HTTP POST method should be used and gives the route path.
// It also associates the middlewares and the route handler.
var loginRoute = route{
	method: http.MethodPost,
	path:   "/login",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired, // Middleware indicating that no authentication is required for this route.
		loginHandler,              // The route handler that will be executed when the route is called.
	},
}

var registrationHandler = func(ctx *octopus.Context) {
	var newUser = models.User{}
	// Attempts to deserialize the form data into the User instance.
	if err := ctx.BodyParser(&newUser); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	err := newUser.Validate()
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while hashing the password.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	newUser.Password = string(newHash)
	// Attempts to create a new user in the database with the provided data.
	if newUser.Create(ctx.Db.Conn) != nil {
		// If user creation fails, logs the error and returns an HTTP status  500.
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
		"message": "User successfully registered and logged.",
		"status":  "200",
	})
}

// registrationRoute is a structure that defines the registration route for the API.
// It specifies that the HTTP POST method should be used and gives the route path.
// It also associates the middlewares and the route handler.
var registrationRoute = route{
	method: http.MethodPost,
	path:   "/registration",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired, // Middleware indicating that no authentication is required for this route.
		registrationHandler,       // The route handler that will be executed when the route is called.
	},
}

func init() {
	// Register the login and registration routes with the global AllHandler map.
	AllHandler[loginRoute.path] = loginRoute
	AllHandler[registrationRoute.path] = registrationRoute
}
