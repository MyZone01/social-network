package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"fmt"
	"net/http"
	"strings"
)

func handleValidSession(ctx *octopus.Context) {
	data := map[string]interface{}{
		"message": "Authenticated successfully",
	}
	ctx.JSON(data)
}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var checkSessionRoute = route{
	path:   "/checksession",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handleValidSession, // Handler function to process the authentication request.
	},
}

func handleUserInfos(ctx *octopus.Context) {
	user := models.User{}

	var token string
	headerBearer := ctx.Request.Header.Get("Authorization")
	if strings.HasPrefix(headerBearer, "Bearer ") {
		token = strings.TrimPrefix(headerBearer, "Bearer ")
	} else {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Invalid Token.",
		})
		return
	}

	id, err := config.Sess.Start(ctx).Get(token)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Invalid Token.",
		})
		return
	}

	erro := user.Get(ctx.Db.Conn, id)
	if (erro != nil) {
		fmt.Println(erro)
	}
	ctx.JSON(user)
}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var userInfosRoute = route{
	path:   "/userinfos",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		handleUserInfos, // Handler function to process the authentication request.
	},
}
