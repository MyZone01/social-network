package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"fmt"
	"net/http"
)

func handleMessages(ctx *octopus.Context) {
	
}
func GetUsers(ctx *octopus.Context) {

	var users models.Users

	err1 := users.GetAll(ctx.Db.Conn)
	if err1 != nil {
		fmt.Println(err1)
		// HandleError(ctx.ResponseWriter, http.StatusInternalServerError, "Error getting users : "+err1.Error())
		return
	}
	fmt.Println(users)

	data := map[string]interface{}{
		"list": users,
	}

	ctx.JSON(data)

	// HandleError(ctx.ResponseWriter, http.StatusUnauthorized, "No active session")

}
func GetMessages(ctx *octopus.Context) {

	var messages models.PrivateMessages

	// err1 := messages.GetPrivateMessages(ctx.Db.Conn, receiverId)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	// HandleError(ctx.ResponseWriter, http.StatusInternalServerError, "Error getting users : "+err1.Error())
	// 	return
	// }

	data := map[string]interface{}{
		"messages_list": messages,
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
	path:   "/users",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.NoAuthRequired, // Middleware to check if the request is authenticated.
		GetUsers,                  // Handler function to process the messages request.
	},
}

func init() {
	// Register the events route with the global AllHandler map.
	AllHandler[getUsers.path] = getUsers
}
