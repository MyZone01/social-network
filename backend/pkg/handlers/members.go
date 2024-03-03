package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func sendInvitationHandler(ctx *octopus.Context) {
	newMember := models.GroupMember{
		Status: models.MemberStatusInvited,
		Role:   models.MemberRoleUser,
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
		UserID:    invitedUserId,
		Type:      models.TypeGroupInvitation,
		Message:   "You have been invited to join a group",
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
		Role:   models.MemberRoleUser,
	}

	group := ctx.Values["group"].(models.Group)
	requestingUserId := ctx.Values["userId"].(uuid.UUID)
	err := newMember.CreateMember(ctx.Db.Conn, requestingUserId, group.ID)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	notification := models.Notification{
		ConcernID: newMember.ID,
		UserID:    group.CreatorID,
		Type:      models.TypeGroupInvitation,
		Message:   "You have been invited to join a group",
	}
	err = notification.Create(ctx.Db.Conn)
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
	group := models.Group{
		ID: ctx.Values["group_id"].(uuid.UUID),
	}
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
	AllHandler[sendInvitationRoute.path] = sendInvitationRoute
	AllHandler[acceptIntegrationRoute.path] = acceptIntegrationRoute
	AllHandler[declineIntegrationRoute.path] = declineIntegrationRoute
	AllHandler[askAccessRoute.path] = askAccessRoute
	AllHandler[getAllRequestAccessRoute.path] = getAllRequestAccessRoute
	AllHandler[acceptRequestAccessRoute.path] = acceptRequestAccessRoute
	AllHandler[declineRequestAccessRoute.path] = declineRequestAccessRoute
}
