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
	Email    string `json:"email"`
	Password string `json:"password"`
}

// loginHandler is a function that handles user login requests.
// It attempts to unmarshal the form data from the client into a User instance,
// checks if the credentials are valid, and if successful, starts a new session for the user.
var loginHandler = func(ctx *octopus.Context) {
	log.Println("Host: [" + ctx.Request.RemoteAddr + "] reach login route")
	var credentials = credentials{}

	if err := ctx.BodyParser(&credentials); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	newUser := models.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	if userCredentialAreValid := newUser.CheckCredentials(ctx); !userCredentialAreValid {
		ctx.Status(http.StatusUnauthorized).JSON(
			map[string]interface{}{
				"error": "credentials are not valid",
			},
		)
		return
	}
	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	if err != nil {
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
		middleware.NoAuthRequired,
		loginHandler,
	},
}

// registrationHandler is a function that handles account creation requests.
// It reads the submitted form data from the client, creates a new user in the database,
// and starts a new session for the user.
var registrationHandler = func(ctx *octopus.Context) {
	log.Println(" Host:  [ " + ctx.Request.RemoteAddr + " ] " + "reach registration route")

	var newUser = models.User{}

	if err := ctx.BodyParser(&newUser); err != nil {
		log.Println(err)
		ctx.Status(500)
		return
	}

	fmt.Println(newUser)

	if err := newUser.Create(ctx.Db.Conn); err != nil {
		log.Println(err)
		ctx.Status(500)
		return
	}

	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	if err != nil {
		log.Println(err)
		ctx.Status(500)
		return
	}
	ctx.JSON(idSession)
}

// registrationRoute is a structure that defines the registration route for the API.
// It specifies that the HTTP POST method should be used and gives the route path.
// It also associates the middlewares and the route handler.
var registrationRoute = route{
	method: http.MethodPost,
	path:   "/registration",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired,
		registrationHandler,
	},
}

var healthHandler = func(ctx *octopus.Context) {
	ctx.WriteString("💻Server is Ok!")
}

var healthRoute = route{
	method: http.MethodGet,
	path: "/health",
	middlewareAndHandler: []octopus.HandlerFunc{
		healthHandler,
	},
}

func init() {
	AllHandler[loginRoute.path] = loginRoute
	AllHandler[registrationRoute.path] = registrationRoute
	AllHandler[healthRoute.path] = healthRoute
}
