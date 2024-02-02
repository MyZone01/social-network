package handlers

import (
	"backend/pkg/models"
	"database/sql"
)

var authHandler = func(r *models.Route) {
	r.
}
var db, err = sql.Open("fkkf", "kdskdskskd")
var authenticationRoute = models.NewRoot("/auth", "GET", true, authHandler, db)

func lolo() {
	authenticationRoute.LunchRoot()
}
