package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"fmt"
	"net/http"
)

// handleMessages is the core function that processes the messages request.
// It receives a Context object containing the request and response writer, along with additional methods for handling the request.
// Use the Context object to implement the messages logic, such as sending or receiving messages.
// After successful operation, you can send a response back to the client using methods like ctx.JSON().
func handleMessages(ctx *octopus.Context) {
	// TODO: Implement the messages logic here.
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

// MessagesHandler defines the structure for handling messages requests.
// It specifies the HTTP method (POST), the path for the endpoint, and the sequence of middleware and handler functions to execute.
var messagesRoutes = route{
	path:   "/messages",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AuthRequired, // Middleware to check if the request is authenticated.
		handleMessages,          // Handler function to process the messages request.
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
