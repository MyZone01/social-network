package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"
)

func insertPostHandler(ctx *octopus.Context) {
	newPost := models.Post{}
	if err := ctx.BodyParser(&newPost); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while  creating new post",
		})
		return
	}
	userPostOwnerId, _ := config.Sess.Start(ctx).Get(ctx.GetBearerToken())
	log.Println(newPost)
	newPost.UserID = userPostOwnerId
	if err := newPost.Create(ctx.Db.Conn); err != nil {
		log.Println(err ,"dfl")
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while creating new post",
		})
		return
	}
	ctx.JSON(newPost.ExploitForRendering(ctx.Db.Conn))
}

func feedHandler(ctx *octopus.Context) {
	log.Println("feedHandler")
	feedPosts := models.Posts{}
	user, _ := config.Sess.Start(ctx).Get(ctx.GetBearerToken())

	if err := feedPosts.GetAvailablePostForUser(ctx.Db.Conn, user); err != nil {
		log.Println(err)
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"error": "error while getting posts",
		})
		return
	}
	ctx.JSON(feedPosts.ExploitForRendering(ctx.Db.Conn))

}

// AuthenticationHandler defines the structure for handling authentication requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var insertPostRoute = route{
	path:   "/post/insert",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		insertPostHandler, // Handler function to process the authentication request.
	},
}

var getFeedPostsRoute = route{
	path:   "/post/getfeed",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		/* ... you can add other middleware here
		   Note: Make sure to place your handler function at the end of the list. */
		feedHandler, // Handler function to process the authentication request.
	},
}

func init() {
	// Register the authentication route with the global AllHandler map.
	AllHandler[getFeedPostsRoute.path] = getFeedPostsRoute
	AllHandler[insertPostRoute.path] = insertPostRoute
}
