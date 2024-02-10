package handlers

import (
	octopus "backend/app" 
	"backend/pkg/config"
	"backend/pkg/middleware" 
	"backend/pkg/models" 
	"io/ioutil" 
	"log" 
	"net/http" 
)

// registationHandler is a function that handles account creation requests.
// It reads the submitted form data from the client, creates a new user in the database,
// and starts a new session for the user.
var registationHandler = func(ctx *octopus.Context) {
	// Logs the client's IP address that reached the registration route.
	log.Println(" Host:  [ " + ctx.Request.RemoteAddr + " ] " + "reach registration route")

	// Reads the HTTP request body and stores the data in jsonForm.
	jsonForm, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// In case of an error reading the request body, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}

	// Creates an empty User instance to store user data.
	var newUser = models.User{}

	// Attempts to deserialize the form data into the User instance.
	if err := newUser.UnmarshalFormData(jsonForm); err != nil {
		// If deserialization fails, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}

	// Attempts to create a new user in the database with the provided data.
	if err := newUser.Create(ctx.Db.Conn); err != nil {
		// If user creation fails, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}

	// Starts a new session for the user and sets the user's ID as the session key.
	if err := config.Sess.Start(ctx).Set(newUser.ID); err != nil {
		// If starting the session fails, logs the error and returns an HTTP status  500.
		log.Println(err)
		ctx.Status(500)
		return
	}

	// Logs the success of the user's registration and login.
	log.Println(" User :" + newUser.Nickname + " with  Host: [ " + ctx.Request.RemoteAddr + " ] " + "is successfully registered and logged")
}

// registrationRoute is a structure that defines the registration route for the API.
// It specifies that the HTTP POST method should be used and gives the route path.
// It also associates the middlewares and the route handler.
var registrationRoute = route{
	method: http.MethodPost,
	path:   "/registrations",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired, // Middleware indicating that no authentication is required for this route.
		registationHandler, // The route handler that will be executed when the route is called.
	},
}
