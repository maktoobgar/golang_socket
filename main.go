package main

import (
	"github.com/gin-gonic/gin"
)

type RoomCreate struct {
	Name              string `json:"name"`
	ConnectionsLength int    `json:"connection_length"`
}

var addr = "127.0.0.1:8080"

var rooms = map[string]*Room{}

func main() {
	r := gin.New()
	r.POST("/room", func(ctx *gin.Context) {
		data := &RoomCreate{}
		if err := ctx.BindJSON(data); err != nil {
			ctx.JSON(400, gin.H{
				"message": "bad body request",
			})
			return
		}

		if room, ok := rooms[data.Name]; ok {
			data.ConnectionsLength = len(room.clients)
			ctx.JSON(200, data)
			return
		}

		room := newRoom()
		go room.run()
		rooms[data.Name] = room
		ctx.JSON(201, data)
	})
	r.GET("/room/:roomName/ws", func(ctx *gin.Context) {
		roomName := ctx.Param("roomName")
		var room *Room = nil
		var ok = false
		if room, ok = rooms[roomName]; !ok {
			return
		}

		connectToRoom(room, ctx.Writer, ctx.Request)
	})
	r.Run(addr)
}
