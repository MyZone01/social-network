// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"net/http"
)

func sendInvitation(ctx *octopus.Context) {
	// TODO: Implement the notifications logic here.
}

var notificationsRoute = route{
	path:   "/invitation",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		sendInvitation,
	},
}

func init() {
	AllHandler[notificationsRoute.path] = notificationsRoute
}
