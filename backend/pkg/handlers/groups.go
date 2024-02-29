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
		log.Println(err)
		return
	}

	newGroup.CreatorID = ctx.Values["userId"].(uuid.UUID)
	if err := newGroup.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newGroup)
}

var groupsCreateRoute = route{
	path:   "/create-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		createGroup,
	},
}

func getGroups(ctx *octopus.Context) {
	groups := models.Groups{}
	isMemberNeeded := ctx.Request.URL.Query().Get("isMemberNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := groups.GetAllGroups(ctx.Db.Conn, isMemberNeeded, isUserNeeded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(groups)
}

func getGroupById(ctx *octopus.Context) {
	group := models.Group{}
	groupID := uuid.MustParse(ctx.Request.URL.Query().Get("id"))
	isMemberNeeded := ctx.Request.URL.Query().Get("isMemberNeeded") == "true"
	isUserNeeded := ctx.Request.URL.Query().Get("isUserNeeded") == "true"
	err := group.Get(ctx.Db.Conn, groupID, isMemberNeeded, isUserNeeded)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(group)
}

func createPostGroup(ctx *octopus.Context) {
	newPost := models.GroupPost{}
	post := models.Post{}
	_groupId := uuid.MustParse(ctx.Request.URL.Query().Get("group_id"))
	newPost.GroupID = _groupId
	newPost.CreatorID = ctx.Values["userId"].(uuid.UUID)

	if err := ctx.BodyParser(&post); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(post, err.Error())
		return
	}

	newPost.Post = post
	newPost.Post.UserID = ctx.Values["userId"].(uuid.UUID)

	if err := newPost.CreatePost(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newPost)
}

func getGroupPosts(ctx *octopus.Context) {
	posts := models.GroupPosts{}
	groupID := uuid.MustParse(ctx.Request.URL.Query().Get("group_id"))
	err := posts.GetPosts(ctx.Db.Conn, groupID, true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(posts)
}

func getGroupPostById(ctx *octopus.Context) {
	post := models.GroupPost{}
	postID, err := uuid.Parse(ctx.Request.URL.Query().Get("post_id"))
	if err != nil{
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": "Invalid post uuid",
		})
		return
	}

	groupId := ctx.Values["group_id"].(uuid.UUID)

	err = post.GetPost(ctx.Db.Conn, groupId, postID, true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(post)
	
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
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getGroupById,
	},
}

var createPostGroupRoute = route{
	path:   "/create-post-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupPostValid,
		createPostGroup,
	},
}

var getGroupPostsRoute = route{
	path:   "/get-all-post-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getGroupPosts,
	},
}

var getGroupPostRoute = route{
	path:   "/get-post-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getGroupPostById,
	},
}

func init() {
	AllHandler[groupsCreateRoute.path] = groupsCreateRoute
	AllHandler[groupsGetRoute.path] = groupsGetRoute
	AllHandler[groupsGetByIdRoute.path] = groupsGetByIdRoute
	AllHandler[createPostGroupRoute.path] = createPostGroupRoute
	AllHandler[getGroupPostsRoute.path] = getGroupPostsRoute
	AllHandler[getGroupPostRoute.path] = getGroupPostRoute
}
