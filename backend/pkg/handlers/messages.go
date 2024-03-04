package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func getAllMessagesByGroup(ctx *octopus.Context) {
	messages := models.GroupMessages{}
	groupId := ctx.Values["group_id"].(uuid.UUID)
	err := messages.GetGroupMessages(ctx.Db.Conn, groupId)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusAccepted).JSON(messages)
}

var getAllMessagesByGroupRoutes = route{
	path:   "/group/messages",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllMessagesByGroup,
	},
}

func addMessageToGroup(ctx *octopus.Context) {
	newMessage := models.GroupMessage{}

	if err := ctx.BodyParser(&newMessage); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newMessage.SenderID = ctx.Values["member"].(*models.GroupMember).ID
	newMessage.GroupID = ctx.Values["group_id"].(uuid.UUID)
	if err := newMessage.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newMessage)
}

var addMessageToGroupRoutes = route{
	path:   "/group/messages/new",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		addMessageToGroup,
	},
}

func init() {
	AllHandler[getAllMessagesByGroupRoutes.path] = getAllMessagesByGroupRoutes
	AllHandler[addMessageToGroupRoutes.path] = addMessageToGroupRoutes
}
