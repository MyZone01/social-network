package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"
)

// loginHandler is a function that handles user login requests.
// It attempts to unmarshal the form data from the client into a User instance,
// checks if the credentials are valid, and if successful, starts a new session for the user.
var loginHandler = func(ctx *octopus.Context) {
	// Log the client's IP address that reached the login route.
	log.Println("Host: [" + ctx.Request.RemoteAddr + "] reach login route")
	var newUser = models.User{}

	// Try to deserialize the form data into the User instance.
	if err := newUser.UnmarshalFormData(ctx); err != nil {
		// If deserialization fails, log the error and return an HTTP status  500.
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
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
	// Start a new session for the user and set the user's ID as the session key.
	if err := config.Sess.Start(ctx).Set(newUser.ID); err != nil {
		// If starting the session fails, log the error and return an HTTP status  500.
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
	}
}

// loginRoute is a structure that defines the login route for the API.
// It specifies that the HTTP POST method should be used and gives the route path.
// It also associates the middlewares and the route handler.
var loginRoute = route{
	method: http.MethodPost,
	path:   "/login",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired, // Middleware indicating that no authentication is required for this route.
		loginHandler,             // The route handler that will be executed when the route is called.
	},
}