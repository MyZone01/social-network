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

var createGroupRoute = route{
	path:   "/create-group",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		createGroup,
	},
}

func getAllGroups(ctx *octopus.Context) {
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

var getAllGroupsRoute = route{
	path:   "/get-all-groups",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		getAllGroups,
	},
}

func getGroupById(ctx *octopus.Context) {
	group := models.Group{}
	groupID := ctx.Values["group_id"].(uuid.UUID)
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

var getGroupByIdRoute = route{
	path:   "/get-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getGroupById,
	},
}

func createPostGroup(ctx *octopus.Context) {
	newPost := models.GroupPost{}
	post := models.Post{}
	newPost.GroupID = ctx.Values["group_id"].(uuid.UUID)
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

func getAllGroupPosts(ctx *octopus.Context) {
	posts := models.GroupPosts{}
	groupID := ctx.Values["group_id"].(uuid.UUID)
	err := posts.GetPosts(ctx.Db.Conn, groupID, true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(posts)
}

var getAllGroupPostsRoute = route{
	path:   "/get-all-post-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllGroupPosts,
	},
}

func getGroupPostById(ctx *octopus.Context) {
	post := models.GroupPost{}
	postID, err := uuid.Parse(ctx.Request.URL.Query().Get("post_id"))
	if err != nil {
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

var getGroupPostRoute = route{
	path:   "/get-post-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.IsGroupPostExist,
		middleware.HaveGroupAccess,
		getGroupPostById,
	},
}

// Group member send invitation handler
func sendInvitationHandler(ctx *octopus.Context) {
	newMember := models.GroupMember{
		Status: models.MemberStatusInvited,
		Role: models.MemberRoleUser,
	}

	groupId := ctx.Values["group_id"].(uuid.UUID)
	invitedUserId := ctx.Values["invited_user_id"].(uuid.UUID)
	err := newMember.CreateMember(ctx.Db.Conn, invitedUserId, groupId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	notification := models.Notification{
		ConcernID: newMember.ID,
		UserID: invitedUserId,
		Type: models.TypeGroupInvitation,
		Message: "You have been invited to join a group",
	}
	err = notification.Create(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(newMember)
}

var sendInvitationRoute = route{
	path:   "/send-invitation",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.HaveGroupAccess,
		middleware.IsGroupExist,
		middleware.IsInvitedUserExist,
		sendInvitationHandler,
	},
}

// Group member accept/decline invitation handler
func acceptIntegrationHandler(ctx *octopus.Context) {
	member := ctx.Values["member"].(models.GroupMember)
	member.Status = models.MemberStatusAccepted
	err := member.UpdateMember(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.Status(http.StatusOK).JSON(member)
}

var acceptIntegrationRoute = route{
	path:   "/accept-invitation",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsInvitationExist,
		acceptIntegrationHandler,
	},
}

func declineIntegrationHandler(ctx *octopus.Context) {
	member := ctx.Values["member"].(models.GroupMember)
	member.Status = models.MemberStatusDeclined
	err := member.UpdateMember(ctx.Db.Conn)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	ctx.Status(http.StatusOK).JSON(member)
}

var declineIntegrationRoute = route{
	path:   "/decline-invitation",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsInvitationExist,
		declineIntegrationHandler,
	},
}

func askAccessHandler(ctx *octopus.Context) {
	newMember := models.GroupMember{
		Status: models.MemberStatusInvited,
		Role: models.MemberRoleUser,
	}

	groupId := ctx.Values["group_id"].(uuid.UUID)
	requestingUserId := ctx.Values["userId"].(uuid.UUID)
	err := newMember.CreateMember(ctx.Db.Conn, requestingUserId, groupId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.JSON(newMember)
}

var askAccessRoute = route{
	path:   "/request-access",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		askAccessHandler,
	},
}

func getAllRequestAccess(ctx *octopus.Context) {
	group := models.Group{}
	err := group.GetMembers(ctx.Db.Conn, models.MemberStatusRequesting, true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	requestingUsers := group.GroupMembers

	ctx.JSON(requestingUsers)
}

var getAllRequestAccessRoute = route{
	path:   "/get-all-request-access",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupAdmin,
		getAllRequestAccess,
	},
}

var acceptRequestAccessRoute = route{
	path:   "/accept-access",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupAdmin,
		middleware.IsRequestExist,
		acceptIntegrationHandler,
	},
}

var declineRequestAccessRoute = route{
	path:   "/accept-decline-access",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		declineIntegrationHandler,
	},
}

func init() {
	AllHandler[createGroupRoute.path] = createGroupRoute
	AllHandler[getAllGroupsRoute.path] = getAllGroupsRoute
	AllHandler[getGroupByIdRoute.path] = getGroupByIdRoute
	AllHandler[createPostGroupRoute.path] = createPostGroupRoute
	AllHandler[getAllGroupPostsRoute.path] = getAllGroupPostsRoute
	AllHandler[getGroupPostRoute.path] = getGroupPostRoute
	AllHandler[sendInvitationRoute.path] = sendInvitationRoute
	AllHandler[acceptIntegrationRoute.path] = acceptIntegrationRoute
	AllHandler[declineIntegrationRoute.path] = declineIntegrationRoute
	AllHandler[askAccessRoute.path] = askAccessRoute
	AllHandler[getAllRequestAccessRoute.path] = getAllRequestAccessRoute
	AllHandler[acceptRequestAccessRoute.path] = acceptRequestAccessRoute
	AllHandler[declineRequestAccessRoute.path] = declineRequestAccessRoute
}
