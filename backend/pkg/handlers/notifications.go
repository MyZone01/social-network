// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
)

func handlernotifications(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	notifications := new(models.Notifications)
	if err := notifications.GetByUser(ctx.Db.Conn, userId); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": err,
		})
		return
	}

	AllNotification := []map[string]interface{}{}
	for _, notification := range *notifications {
		user := new(models.User)
		user.Get(ctx.Db.Conn, notification.UserID)
		user.Password = ""
		AllNotification = append(AllNotification, map[string]interface{}{
			"type":       notification.Type,
			"concernID":  notification.ConcernID,
			"user":       user,
			"message":    notification.Message,
			"created_at": notification.CreatedAt,
		})
	}
	ctx.JSON(AllNotification)
}

var notificationsRoute = route{
	path:   "/getnotifications",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handlernotifications,
	},
}

func init() {
	AllHandler[notificationsRoute.path] = notificationsRoute
}
