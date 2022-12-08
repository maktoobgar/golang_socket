package rooms

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/chat"
)

func Delete(ctx *gin.Context) {
	var (
		roomName            = ctx.Param("roomName")
		room     *chat.Room = nil
		ok                  = false
	)
	if room, ok = chat.Rooms[roomName]; !ok {
		ctx.Status(400)
		return
	}

	delete(chat.Rooms, roomName)
	room.Terminate()
	ctx.Status(204)
}
