package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/handlers/ws"
)

func joinRooms(r *gin.Engine) {
	r.GET("/rooms/:roomName/ws", ws.JoinRooms)
}
