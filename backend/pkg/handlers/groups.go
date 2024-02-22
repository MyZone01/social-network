// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func createGroup(ctx *octopus.Context) {
	newGroup := models.Group{}

	if err := ctx.BodyParser(&newGroup); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	newGroup.CreatorID = ctx.Values["userId"].(uuid.UUID)
	if err := newGroup.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newGroup)
}

func getGroups(ctx *octopus.Context) {
	groups := models.Groups{}
	isMemberNeeded := ctx.Request.URL.Query().Get("isMemberNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := groups.GetAllGroups(ctx.Db.Conn, isMemberNeeded, isUserNeeded)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(groups)
}

// Get group by id
func getGroupById(ctx *octopus.Context) {
	group := models.Group{}
	groupID := uuid.MustParse(ctx.Request.URL.Query().Get("id"))
	isMemberNeeded := ctx.Request.URL.Query().Get("isMemberNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := group.Get(ctx.Db.Conn, groupID, isMemberNeeded, isUserNeeded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.JSON(group)
}

func createPostGroup(ctx *octopus.Context) {
	newPost := models.GroupPost{}
	newPost.GroupID = uuid.MustParse(ctx.Request.URL.Query().Get("group_id"))
	newPost.CreatorID = ctx.Values["userId"].(uuid.UUID)

	if err := ctx.BodyParser(&newPost); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	if err := newPost.CreatePost(ctx.Db.Conn, ); err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newPost)
}

var groupsCreateRoute = route{
	path:   "/create-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		createGroup,
	},
}

var groupsGetRoute = route{
	path:   "/get-all-groups",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		getGroups,
	},
}

var groupsGetByIdRoute = route{
	path:   "/get-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		getGroupById,
	},
}

// create post group
var createPostGroupRoute = route{
	path:   "/create-post-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		createPostGroup,
	},
}

func init() {
	AllHandler[groupsCreateRoute.path] = groupsCreateRoute
	AllHandler[groupsGetRoute.path] = groupsGetRoute
	AllHandler[groupsGetByIdRoute.path] = groupsGetByIdRoute
	AllHandler[createPostGroupRoute.path] = createPostGroupRoute
}
