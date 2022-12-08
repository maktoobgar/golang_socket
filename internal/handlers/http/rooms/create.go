package rooms

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/chat"
	"github.com/maktoobgar/golang_socket/internal/dto"
)

func Create(ctx *gin.Context) {
	data := &dto.Room{}
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad body request",
		})
		return
	}

	if room, ok := chat.Rooms[data.Name]; ok {
		data.ConnectionsLength = len(room.Clients)
		ctx.JSON(200, data)
		return
	}

	room := chat.NewRoom()
	go room.Run()
	chat.Rooms[data.Name] = room
	ctx.JSON(201, data)
}
