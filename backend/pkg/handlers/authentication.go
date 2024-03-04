package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"errors"
	"fmt"
	"log"
	"net/http"

	"net/mail"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *credentials) Validate() error {
	if _, err := mail.ParseAddress(c.Email); err != nil {
		return errors.New("Invalid email")
	}

	if c.Password == "" {
		return errors.New("Password is missing. Please provide it.")
	}

	return nil
}

var loginHandler = func(ctx *octopus.Context) {
	var credentials = credentials{}

	if err := ctx.BodyParser(&credentials); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while parsing the form data.",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	if err := credentials.Validate(); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	newUser := models.User{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	err := newUser.Get(ctx.Db.Conn, credentials.Email)
	if err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "invalid email.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	fmt.Println(newUser.Password, bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(credentials.Password)))

	if err := bcrypt.CompareHashAndPassword([]byte(newUser.Password), []byte(credentials.Password)); err != nil {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]interface{}{
			"session": "",
			"message": "Invalid credentials. Please try again.",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while starting the session.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.JSON(map[string]interface{}{
		"session": idSession,
		"user":    newUser,
		"message": "User successfully logged.",
		"status":  "200",
	})
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
	var newUser = models.User{}
	if err := ctx.BodyParser(&newUser); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	err := newUser.Validate()
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"session": "",
			"message": err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while hashing the password.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	newUser.Password = string(newHash)
	if err := newUser.Create(ctx.Db.Conn); err != nil {
		log.Println(err.Error())
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while creating the user.",
			"status":  http.StatusInternalServerError,
		})
		return
	}

	idSession, err := config.Sess.Start(ctx).Set(newUser.ID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"session": "",
			"message": "Error while starting the session.",
			"status":  http.StatusInternalServerError,
		})
		return
	}
	ctx.Status(http.StatusAccepted).JSON(map[string]interface{}{
		"session": idSession,
		"user":    newUser,
		"message": "User successfully registered and logged.",
		"status":  "200",
	})
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
	path:   "/health",
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

func meHandler(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)
	user := models.User{}
	err := user.Get(ctx.Db.Conn, userId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.Status(http.StatusOK).JSON(user)
}

var meRoute = route{
	method: http.MethodGet,
	path:   "/me",
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		meHandler,
	},
}

func init() {
	AllHandler[loginRoute.path] = loginRoute
	AllHandler[logoutRoute.path] = logoutRoute
	AllHandler[meRoute.path] = meRoute
	AllHandler[registrationRoute.path] = registrationRoute
	AllHandler[healthRoute.path] = healthRoute
}
