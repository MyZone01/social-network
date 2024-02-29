package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func createPost(ctx *octopus.Context) {
	newPost := models.Post{}
	if err := ctx.BodyParser(&newPost); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(newPost, err)
	}

	_userID := ctx.Values["userId"].(uuid.UUID)

	if err := newPost.Create(ctx.Db.Conn, _userID); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newPost)
}

var postRoute = route{
	path:   "/create-post",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsPostValid,
		createPost,
	},
}

func init() {
	AllHandler[postRoute.path] = postRoute
}
