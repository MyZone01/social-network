package handlers

import (
	octopus "backend/app"
	"net/http"
)

// HandlerConstructor is a type alias for a function that takes a path and a variadic number of HandlerFuncs.
// It's used to define the constructor for creating new routes with associated middleware and handlers.
type HandlerConstructor func(path string, middlewareAndHandler ...octopus.HandlerFunc)

// Handler represents a route with its associated path, constructor, and middleware/handler functions.
type Handler struct {
	path, method         string
	middlewareAndHandler []octopus.HandlerFunc
}

// AllHandler is a slice of Handler structures that defines all the routes for the application.
// Each Handler in the slice includes the path, the constructor for creating the route, and the middleware/handler functions to be executed.
var AllHandler = []Handler{
	authenticationHandler,
	MessagesHandler,
	postHandler,
	GroupsHandler,
	NotificationsHandler,
	EventsHandler,
	EventsHandler,
	// Add more handlers here as needed.
}

// HandleAll is a function that iterates over the AllHandler slice and applies each Handler's constructor to register the routes.
// This function should be called during the initialization phase of the application to set up all the routes.
var HandleAll = func(app *octopus.App) {
	var mapContructors = map[string]HandlerConstructor{
		http.MethodGet:  app.GET,
		http.MethodPost: app.POST,
	}
	for _, v := range AllHandler {
		mapContructors[v.method](v.path, v.middlewareAndHandler...)
	}
}
