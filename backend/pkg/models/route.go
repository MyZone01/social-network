package models

import (
	"database/sql"
	"net/http"
)

// Route represents an HTTP route handler.
type Route struct {
	R            *http.Request       // HTTP request
	W            http.ResponseWriter // HTTP response writer
	db           *sql.DB             // Database connection
	path         string              // Path for the route
	finalHandler func(w http.ResponseWriter, r *http.Request)
}

// LunchRoot registers the route and its final handler with the HTTP server.
func (r *Route) LunchRoot() {
	http.HandleFunc(r.path, r.finalHandler)
}

// GetUserInfo retrieves user information.
func (r *Route) GetUserInfo() {
	// Implementation for retrieving user information
	// Placeholder comment, awaiting database implementation.
}

// SelectData performs a database SELECT operation.
func (r *Route) SelectData() {
	// Implementation for selecting data from the database
	// Placeholder comment, awaiting database implementation.
}

// InsertData performs a database INSERT operation.
func (r *Route) InsertData() {
	// Implementation for inserting data into the database
	// Placeholder comment, awaiting database implementation.
}

func (r *Route) UpdateData() {

}

func (r *Route) DeleteData() {

}

// Checkauth is a sample authentication function.
var Checkauth = func(r *http.Request) bool {
	// Placeholder authentication logic
	// Placeholder comment, awaiting database implementation.
	return true
}

// NewRoot creates a new Route with the specified parameters.
func NewRoot(path, method string, authVerification bool, customHandler func(r *Route), db *sql.DB) Route {
	rootToReturn := Route{path: path}

	// finalFunc is the final handler function for the route.
	finalFunc := func(w http.ResponseWriter, r *http.Request) {
		// DefaultMiddleware : set of middleware checks applied to all routes
		// - Check if the HTTP method matches the expected method.
		// - Check if the requested URL path matches the expected path.
		// - Optionally, perform authentication check using Checkauth function.
		if r.Method != method {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if r.URL.Path != path {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if authVerification && Checkauth(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		rootToReturn.R = r
		rootToReturn.W = w
		rootToReturn.db = db

		// Call the customHandler with the Route instance.
		customHandler(&rootToReturn)
	}

	// Set the final handler for the route.
	rootToReturn.finalHandler = finalFunc

	return rootToReturn
}
