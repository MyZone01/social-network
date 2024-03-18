package handlers

import (
	octopus "backend/app"
	"backend/pkg/middleware"
	"backend/pkg/models"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conns = new(sync.Map)
)

func handleSocket(ctx *octopus.Context) {
	conn, err := upgrader.Upgrade(ctx.ResponseWriter, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conns.Store(uuid.New(), conn)
	for {
		go func() {
			for {
				models.Data.Range(func(kk, value interface{}) bool {
					key, ok := kk.(string)
					val, okval := value.(map[string]interface{})
					if ok && okval {
						conns.Range(func(k, v interface{}) bool {
							if v, ok := v.(*websocket.Conn); ok {
								err := v.WriteJSON(map[string]interface{}{
									"data": val,
									"type": key,
								})
								if err != nil {
									conns.Delete(k)
									if err := v.Close(); err != nil {
										log.Println(err)
									}
								}
							}
							return true
						})
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
			err := conn.WriteJSON(map[string]interface{}{
				"status":  http.StatusBadRequest,
				"message": "Invalid message",
			})
			if err != nil {
				if err := conn.Close(); err == nil {
					log.Println(err)
				}
			}
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
				err := conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusBadRequest,
					"message": "Invalid message",
				})
				if err != nil {
					if err := conn.Close(); err == nil {
						log.Println(err)
					}
				}
				return
			}
			user := models.User{}
			if user.Get(ctx.Db.Conn, newMesssage.ReceiverID) != nil {
				err := conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusNotFound,
					"message": "User not found",
				})
				if err != nil {
					if err := conn.Close(); err == nil {
						log.Println(err)
					}
				}
				return
			}
			if user.Get(ctx.Db.Conn, newMesssage.SenderID) != nil {
				err := conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusNotFound,
					"message": "User not found",
				})
				if err != nil {
					if err := conn.Close(); err == nil {
						log.Println(err)
					}
				}
				return
			}
			if err := newMesssage.Create(ctx.Db.Conn); err != nil {
				err = conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Internal server error",
				})
				if err != nil {
					if err := conn.Close(); err == nil {
						log.Println(err)
					}
				}
				return
			}
			newNotification := models.Notification{
				UserID:    newMesssage.SenderID,
				ConcernID: newMesssage.ReceiverID,
				Type:      models.TypeNewMessage,
				Message:   newMesssage.Content,
			}
			if err := newNotification.Create(ctx.Db.Conn); err != nil {
				err = conn.WriteJSON(map[string]interface{}{
					"status":  http.StatusInternalServerError,
					"message": "Internal server error",
				})
				if err != nil {
					if err := conn.Close(); err == nil {
						log.Println(err)
					}
				}
				return
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
