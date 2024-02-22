package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"net/http"
)

func handlePost(ctx *octopus.Context) {

}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var postRoute = route{
	path:   "/post",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handlePost, // Handler function to process the authentication request.
	},
}

func init() {
	// Register the authentication route with the global AllHandler map.
	AllHandler[postRoute.path] = postRoute
}
