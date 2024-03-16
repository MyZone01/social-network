package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func handleSocket(ctx *octopus.Context) {
	conn, err := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		go func() {
			for {
				models.Data.Range(func(key, value interface{}) bool {
					if err := conn.WriteJSON(map[string]interface{}{
						"data": value,
						"type": strings.Split(key.(string), "_id_")[0],
					}); err != nil {
						log.Println(err)
						return false
					}
					models.Data.Delete(key)
					return true
				})
				time.Sleep(time.Second)
			}
		}()
		var data = map[string]interface{}{}
		// newMesssage := models.PrivateMessage{}
		err := conn.ReadJSON(&data)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(data)

		if data["type"] == "private_message" {
			newMesssage := new(models.PrivateMessage)
			msg := data["message"].(map[string]interface{})
			newMesssage.Content = msg["content"].(string)
			newMesssage.SenderID = uuid.MustParse(msg["sender_id"].(string))
			newMesssage.ReceiverID = uuid.MustParse(msg["receiver_id"].(string))
			if newMesssage.Content == "" || newMesssage.SenderID == uuid.Nil || newMesssage.ReceiverID == uuid.Nil {
				conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusBadRequest,
					"message": "Invalid message",
				})
				return
			}
			user := models.User{}
			if user.Get(ctx.Db.Conn, newMesssage.ReceiverID) != nil {
				conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusNotFound,
					"message": "User not found",
				})
				return
			}
			if user.Get(ctx.Db.Conn, newMesssage.SenderID) != nil {
				conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusNotFound,
					"message": "User not found",
				})
				return
			}
			if err := newMesssage.Create(ctx.Db.Conn); err != nil {
				conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Internal server error",
				})
			}
			newNotification := models.Notification{
				UserID:    newMesssage.SenderID,
				ConcernID: newMesssage.ReceiverID,
				Type:      models.TypeNewMessage,
				Message:   newMesssage.Content,
			}
			if err := newNotification.Create(ctx.Db.Conn); err != nil {
				conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Internal server error",
				})
			}
		} else if data["type"] == "group_message" {
			// newMesssage, ok := data["message"].(models.GroupMessage)
			// if !ok {
			// 	conn.WriteJSON(map[string]interface{}{
			// 		"status":  http.StatusBadRequest,
			// 		"message": "Invalid message",
			// 	})
			// }

			// if err := newMesssage.Create(ctx.Db.Conn); err != nil {
			// 	conn.WriteJSON(map[string]interface{}{
			// 		"status":  http.StatusInternalServerError,
			// 		"message": "Internal server error",
			// 	})
		}

		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}

var handleSocketRoute = route{
	path:   "/socket",
	method: http.MethodGet,
	middlewareAndHandler: []octopus.HandlerFunc{
		middleware.AllowedSever,
		handleSocket, // Handler function to process the authentication request.
	},
}

func init() {
	// Register the authentication route with the global AllHandler map.
	AllHandler[handleSocketRoute.path] = handleSocketRoute
}
