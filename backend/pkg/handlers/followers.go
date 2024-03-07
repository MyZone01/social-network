package handlers

import (
	octopus "backend/app"
	"backend/pkg/config"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
)

func handleFollower(ctx *octopus.Context) {
	userId := ctx.Values["userId"].(uuid.UUID)

	type request struct {
		Action string `json:"action"`
		Id     string `json:"nickname"`
	}

	req := new(request)
	if err := ctx.BodyParser(req); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Error while parsing the form data.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	user := new(models.User)

	if err := user.Get(ctx.Db.Conn, req.Id); err != nil {
		ctx.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"message": "User not found.",
			"status":  http.StatusNotFound,
		})
		return
	}

	if user.ID == userId {
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid request.",
			"status":  http.StatusBadRequest,
		})
		return
	}

	follow := new(models.Follower)
	follow.FollowerID = userId
	follow.FolloweeID = user.ID

	var reverse = false
	if req.Action == "accept" || req.Action == "decline" {
		reverse = true
	}
	follow.Get(ctx.Db.Conn, reverse)

	notif := new(models.Notification)

	switch req.Action {
	case "follow":
		if follow.Status == models.StatusRequested || follow.Status == models.StatusAccepted || follow.Status == models.StatusDeclined {
			ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
				"message": "You have already sent a follow request.",
				"status":  http.StatusBadRequest,
			})
			return
		}
		if err := follow.Create(ctx.Db.Conn); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
		notif.UserID = follow.FolloweeID
		notif.Type = models.TypeFollowRequest
		notif.Message = user.FirstName + " " + user.LastName + " " + "wants to follow you"
		createNotif(notif, ctx)

		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "Follow request sent successfully",
			"status":  http.StatusOK,
		})
		return
	case "unfollow":
		if follow.Status != models.StatusAccepted && follow.Status != models.StatusRequested {
			ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
				"message": "You are not following this user.",
				"status":  http.StatusBadRequest,
			})
			return

		}
		if err := follow.Delete(ctx.Db.Conn); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
		notif.UserID = follow.FolloweeID
		notif.Type = models.TypeUnFollow
		notif.Message = user.FirstName + " " + user.LastName + " " + "unfollow you"
		createNotif(notif, ctx)

		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "Unfollowed successfully",
			"status":  http.StatusOK,
		})
		return
	case "accept":

		if follow.Status != models.StatusRequested {
			ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
				"message": "Invalid request.",
				"status":  http.StatusBadRequest,
			})
			return
		}
		follow.Status = models.StatusAccepted
		if err := follow.Update(ctx.Db.Conn); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
		notif.UserID = follow.FollowerID
		notif.Type = models.TypeFollowAccepted
		notif.Message = user.FirstName + " " + user.LastName + " " + "accepted your follow request"
		createNotif(notif, ctx)

		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "Request accepted successfully",
			"status":  http.StatusOK,
		})
		return
	case "decline":
		if follow.Status != models.StatusRequested {
			ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
				"message": "Invalid request.",
				"status":  http.StatusBadRequest,
			})
			return
		}
		follow.Status = models.StatusDeclined
		if err := follow.Update(ctx.Db.Conn); err != nil {
			ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
				"message": err.Error(),
				"status":  http.StatusInternalServerError,
			})
			return
		}
		notif.UserID = follow.FollowerID
		notif.Type = models.TypeFollowDeclined
		notif.Message = user.FirstName + " " + user.LastName + " " + "declined your follow request"
		createNotif(notif, ctx)
		ctx.Status(http.StatusOK).JSON(map[string]interface{}{
			"message": "Request declined successfully",
			"status":  http.StatusOK,
		})
		return
	default:
		ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"message": "Invalid action.",
			"status":  http.StatusBadRequest,
		})
		return
	}
}

func createNotif(notif *models.Notification, ctx *octopus.Context) {
	allType := []models.NotificationType{models.TypeFollowAccepted, models.TypeFollowDeclined, models.TypeFollowRequest, models.TypeUnFollow}
	for _, t := range allType {
		newNotif := new(models.Notification)
		newNotif.Type = t
		newNotif.UserID = notif.UserID
		err := notif.Get(ctx.Db.Conn)
		if err != nil {
			continue
		}
		err = newNotif.Delete(ctx.Db.Conn)
		if err != nil {
			continue
		}
	}
	if err := notif.Create(ctx.Db.Conn); err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(map[string]interface{}{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}
	return
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
