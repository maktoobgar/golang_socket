package rooms

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/chat"
	"github.com/maktoobgar/golang_socket/internal/dto"
)

var (
	i   = 0
	key = &sync.Mutex{}
)

func Create(ctx *gin.Context) {
	key.Lock()
	data := &dto.Room{}
	if err := ctx.BindJSON(data); err != nil {
		ctx.JSON(400, gin.H{
			"message": "bad body request",
		})
		return
	}

	if room, ok := chat.Rooms[data.Name]; ok {
		data.ConnectionsLength = len(room.Clients)
		data.ID = room.ID
		ctx.JSON(200, data)
		return
	}

	room := chat.NewRoom(data.Name, i)
	data.ID = i
	i++
	go room.Run()
	chat.Rooms[data.Name] = room
	ctx.JSON(201, data)
	key.Unlock()
}
