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

var loginRoute = route{
	method: http.MethodPost,
	path:   "/login",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired,
		loginHandler,
	},
}

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

func LogoutHandler(ctx *octopus.Context) {
	token := ctx.Values["token"].(string)
	err := config.Sess.Start(ctx).Delete(token)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": "Error while deleting session",
		})
		log.Println(err)
		return
	}

	ctx.Status(http.StatusOK)
}

var logoutRoute = route{
	method: http.MethodDelete,
	path:   "/logout",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		LogoutHandler,
	},
}

func init() {
	AllHandler[loginRoute.path] = loginRoute
	AllHandler[logoutRoute.path] = logoutRoute
	AllHandler[registrationRoute.path] = registrationRoute
	AllHandler[healthRoute.path] = healthRoute
}
