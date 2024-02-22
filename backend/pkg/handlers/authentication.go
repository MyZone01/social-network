package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"fmt"
	"log"
	"net/http"
)

type credentials struct {
	email    string `json:"email"`
	password string `json:"password"`
}

// loginHandler is a function that handles user login requests.
// It attempts to unmarshal the form data from the client into a User instance,
// checks if the credentials are valid, and if successful, starts a new session for the user.
var loginHandler = func(ctx *octopus.Context) {
	// Log the client's IP address that reached the login route.
	log.Println("Host: [" + ctx.Request.RemoteAddr + "] reach login route")
	var credentials = credentials{}

	// Try to deserialize the form data into the User instance.
	if err := ctx.BodyParser(&credentials); err != nil {
		// If deserialization fails, log the error and return an HTTP status  500.
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Email:    credentials.email,
		Password: credentials.password,
	}
	// Check if the user's credentials are valid.
	if userCredentialAreValid := newUser.CheckCredentials(ctx); !userCredentialAreValid {
		// If the credentials are not valid, return an HTTP status  401 with an error message.
		ctx.Status(http.StatusUnauthorized).JSON(
			map[string]interface{}{
				"error": "credentials are not valid",
			},
		)
		return
	}
	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	// Start a new session for the user and set the user's ID as the session key.
	if err != nil {
		// If starting the session fails, log the error and return an HTTP status  500.
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
	ctx.JSON(idSession)
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

// registrationHandler is a function that handles account creation requests.
// It reads the submitted form data from the client, creates a new user in the database,
// and starts a new session for the user.
var registrationHandler = func(ctx *octopus.Context) {
	// Logs the client's IP address that reached the registration route.
	log.Println(" Host:  [ " + ctx.Request.RemoteAddr + " ] " + "reach registration route")

	var newUser = models.User{}

	// Attempts to deserialize the form data into the User instance.
	if err := ctx.BodyParser(&newUser); err != nil {
		// If deserialization fails, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}

	fmt.Println(newUser)

	// Attempts to create a new user in the database with the provided data.
	if err := newUser.Create(ctx.Db.Conn); err != nil {
		// If user creation fails, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}

	// Starts a new session for the user and sets the user's ID as the session key.
	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	if err != nil {
		// If starting the session fails, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}
	ctx.JSON(idSession)
	// Logs the success of the user's registration and login.
	// log.Println(" User :" + newUser.Nickname + " with  Host: [ " + ctx.Request.RemoteAddr + " ] " + "is successfully registered and logged")
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
