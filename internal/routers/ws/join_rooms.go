package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/handlers/ws"
)

func joinRooms(r *gin.Engine) {
	r.GET("/rooms/:roomName/name/:name/ws", ws.JoinRooms)
}
