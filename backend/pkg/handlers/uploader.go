package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"fmt"
	"net/http"
)

func handleAvatarUpload(ctx *octopus.Context) {
	url := ctx.Values["file"]
	fmt.Println(url)

	data := map[string]interface{}{
		"imageurl": url,
	}
	ctx.JSON(data)
}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var avatarUploadRoute = route{
	path:   "/avatarupload",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		middleware.ImageUploadMiddleware,
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handleAvatarUpload, // Handler function to process the authentication request.
	},
}