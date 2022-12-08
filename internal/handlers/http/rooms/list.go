package rooms

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/chat"
	"github.com/maktoobgar/golang_socket/internal/dto"
)

func List(ctx *gin.Context) {
	output := make([]dto.Room, 0, len(chat.Rooms))

	for _, v := range chat.Rooms {
		element := dto.Room{
			ID:                v.ID,
			Name:              v.Name,
			ConnectionsLength: len(v.Clients),
		}
		output = append(output, element)
	}

	ctx.JSON(200, output)
}
