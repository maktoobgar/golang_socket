package rooms

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/chat"
	"github.com/maktoobgar/golang_socket/internal/dto"
)

func List(ctx *gin.Context) {
	keys := make([]dto.Room, len(chat.Rooms))

	i := 0
	for k, v := range chat.Rooms {
		keys[i].Name = k
		keys[i].ConnectionsLength = len(v.Clients)
		i++
	}

	ctx.JSON(200, keys)
}
