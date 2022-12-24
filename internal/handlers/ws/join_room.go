package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/chat"
)

func JoinRooms(ctx *gin.Context) {
	var (
		roomName            = ctx.Param("roomName")
		name                = ctx.Param("name")
		room     *chat.Room = nil
		ok                  = false
	)
	if room, ok = chat.Rooms[roomName]; !ok {
		ctx.Status(400)
		return
	}

	chat.ConnectToRoom(room, ctx.Writer, ctx.Request, name)
}
