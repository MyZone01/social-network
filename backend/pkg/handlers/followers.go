package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"net/http"
)

// handleFollower is the core function that processes the follower request.
// It receives a Context object containing the request and response writer, along with additional methods for handling the request.
// Use the Context object to implement the follower logic, such as adding or removing followers.
// After successful operation, you can send a response back to the client using methods like ctx.JSON().
func handleFollower(ctx *octopus.Context) {
	// TODO: Implement the follower logic here.
}

// FollowerHandler defines the structure for handling follower requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var FollowerRoute = route{
	path:   "/follower",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleFollower,          // Handler function to process the follower request.
	},
}
