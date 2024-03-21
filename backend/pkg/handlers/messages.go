package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"net/http"

	"github.com/google/uuid"
)

func handleMessages(ctx *octopus.Context) {
}

func GetUsers(ctx *octopus.Context) {
	var users models.Users
	var userId = ctx.Values["userId"].(uuid.UUID)

	err1 := users.GetFlow(ctx.Db.Conn, userId)
	if err1 != nil {
		// HandleError(ctx.ResponseWriter, http.StatusInternalServerError, "Error getting users : "+err1.Error())
		return
	}
	data := map[string]interface{}{
		"status": http.StatusOK,
		"data":   users,
	}
	ctx.JSON(data)
	// HandleError(ctx.ResponseWriter, http.StatusUnauthorized, "No active session")
}
func handlerGetMessages(ctx *octopus.Context) {
	var senderId = ctx.Values["userId"].(uuid.UUID)
	var messages models.PrivateMessages
	var receiverId map[string]interface{}
	if err := ctx.BodyParser(&receiverId); err != nil {
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{
			"status":  "400",
			"message": "bad request"})
		return
	}
	err1 := messages.GetPrivateMessages(ctx.Db.Conn, uuid.MustParse(receiverId["receiver_id"].(string)), senderId)
	if err1 != nil {
		// HandleError(ctx.ResponseWriter, http.StatusInternalServerError, "Error getting users : "+err1.Error())
		ctx.Status(http.StatusBadRequest).JSON(map[string]string{"message": "bad request"})
		return
	}
	data := map[string]interface{}{
		"status": http.StatusOK,
		"data":   messages,
	}
	ctx.JSON(data)
	// HandleError(ctx.ResponseWriter, http.StatusUnauthorized, "No active session")
}

var messagesRoutes = route{
	path:   "/groups/messages",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired,
		handleMessages,
	},
}
var getUsers = route{
	path:   "/chatlist",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		GetUsers,                // Handler function to process the messages request.
	},
}
var getMessages = route{
	path:   "/getMessages",
	method: http.MethodPost,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handlerGetMessages,      // Handler function to process the messages request.
	},
}

func init() {
	// Register the events route with the global AllHandler map.
	AllHandler[getUsers.path] = getUsers
	AllHandler[getMessages.path] = getMessages
}
