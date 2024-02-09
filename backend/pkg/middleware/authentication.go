package middleware

import (
	octopus "backend/app"
	"backend/pkg/config"
	"net/http"
)

// AuthRequired verify if the user is connected
func AuthRequired(ctx *octopus.Context) {
	if !config.Sess.Start(ctx).Valid() {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous n'êtes pas connecté.",
		})
		return
	}
	ctx.Next()
}

// NoAuthRequired verify if the user is not connected
func NoAuthRequired(ctx *octopus.Context) {
	if config.Sess.Start(ctx).Valid() {
		ctx.Status(http.StatusUnauthorized).JSON(map[string]string{
			"error": "Vous êtes déja connecté.",
		})
		return
	}
	ctx.Next()
}
