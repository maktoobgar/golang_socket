package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/handlers/http/rooms"
)

func roomsResource(r *gin.Engine) {
	r.GET("/rooms", rooms.List)
	r.POST("/rooms", rooms.Create)
	r.DELETE("/rooms/:roomName", rooms.Delete)
}
