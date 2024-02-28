package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"
)

// handleFollower is the core function that processes the follower request.
// It receives a Context object containing the request and response writer, along with additional methods for handling the request.
// Use the Context object to implement the follower logic, such as adding or removing followers.
// After successful operation, you can send a response back to the client using methods like ctx.JSON().
func handleFollower(ctx *octopus.Context) {
	// TODO: Implement the follower logic here.
}
func handleGetAllFollowersRequest(ctx *octopus.Context) {
	userUUID, err := config.Sess.Start(ctx).Get(ctx.GetBearerToken())
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	userFolowers := models.Followers{}
	userFolowers.GetAllByFolloweeID(ctx.Db.Conn, userUUID)
	userFolowersJson := []map[string]interface{}{}
	for _, follower := range userFolowers {
		newUser := models.User{}
		if err := newUser.Get(ctx.Db.Conn, follower.FollowerID); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		userFolowersJson = append(userFolowersJson,
			map[string]interface{}{
				"nickname":  newUser.Nickname,
				"email":     newUser.Email,
				"firstname": newUser.FirstName,
				"lastname":  newUser.LastName,
				"id":        newUser.ID.String(),
			},
		)
	}
	ctx.JSON(userFolowersJson)

}

// FollowerHandler defines the structure for handling follower requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var FollowerRoute = route{
	path:   "/follower",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleFollower,          // Handler function to process the follower request.
	},
}

var getAllFollowers = route{
	path:   "/getAllFollowers",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,      // Middleware to check if the request is authenticated.
		handleGetAllFollowersRequest, // Handler function to process the follower request.
	},
}

func init() {
	// Register the follower route with the global AllHandler map.
	AllHandler[FollowerRoute.path] = FollowerRoute
	AllHandler[getAllFollowers.path] = getAllFollowers
}
