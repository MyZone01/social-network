package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"net/http"
)

func handleAuthentication(ctx *octopus.Context) {
	ctx.WriteString("hello world")
	ctx.Db.User.Create(ctx.Db.Conn)
	/* The 'ctx' variable contains the same arguments as a typical handler function with 'r' and 'w'.
	and it have new methods such as 'ctx.JSON(data interface{})' to send JSON responses and stuff.
	*** Example usage:
		--- retrieving r and w
			r := ctx.Request
			w := ctx.ResponseWriter

		--- Send a JSON response
			data := map[string]interface{}{
				"message": "Authenticated successfully",
			}
			ctx.JSON(data)

	*** TODO: Put the logic authentication here ***
	*/
}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var authenticationHandler = route{
	path:   "/lolo",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthMiddleware, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handleAuthentication, // Handler function to process the authentication request.
	},
}
