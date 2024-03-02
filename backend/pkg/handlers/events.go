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

func createEventHandler(ctx *octopus.Context) {
	newEvent := models.Event{}

	if err := ctx.BodyParser(&newEvent); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	newEvent.CreatorID = ctx.Values["userId"].(uuid.UUID)
	newEvent.GroupID = ctx.Values["group_id"].(uuid.UUID)
	if err := newEvent.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusCreated).JSON(newEvent)
}

var createEventRoute = route{
	path:   "/create-event",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		createEventHandler,
	},
}

func getAllEventByGroup(ctx *octopus.Context) {
	events := models.Events{}
	groupId := ctx.Values["group_id"].(uuid.UUID)
	err := events.GetGroupEvents(ctx.Db.Conn, groupId, true, true)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	ctx.Status(http.StatusAccepted).JSON(events)
}

var getAllEventRoute = route{
	path:   "/get-all-event-group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		getAllEventByGroup,
	},
}

func respondEventHandler(ctx *octopus.Context) {
	event := ctx.Values["event"].(models.Event)
	member := ctx.Values["member"].(models.GroupMember)
	participant := models.EventParticipant{}
	_participant := models.EventParticipant{}
	if err := ctx.BodyParser(&_participant); err != nil {
		ctx.Status(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err := participant.GetParticipant(ctx.Db.Conn, event.ID, member.ID, false)
	participant.Response = _participant.Response
	if err != nil {
		err := participant.CreateParticipant(ctx.Db.Conn, event.ID, member.ID)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	} else {
		err := participant.UpdateParticipant(ctx.Db.Conn)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}



	ctx.Status(http.StatusOK).JSON(participant)
}

// TODO: Check if the event is not passed
var respondEventRoute = route{
	path:   "/response-event",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		middleware.IsGroupExist,
		middleware.HaveGroupAccess,
		middleware.IsGroupAdmin,
		middleware.IsEventExist,
		respondEventHandler,
	},
}

func init() {
	AllHandler[createEventRoute.path] = createEventRoute
	AllHandler[getAllEventRoute.path] = getAllEventRoute
	AllHandler[respondEventRoute.path] = respondEventRoute
}
