// Package handlers contains the handler functions for various routes.
package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"net/http"
)

func handleGroups(ctx *octopus.Context) {
	ctx.WriteString("Hello, groups!")
}

var groupsRoute = route{
	path:   "/group",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handleGroups,
	},
}
