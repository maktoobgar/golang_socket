package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/core"
)

type RoomCreate struct {
	Name              string `json:"name"`
	ConnectionsLength int    `json:"connection_length"`
}

var addr = "127.0.0.1:8080"

var rooms = map[string]*core.Room{}

func main() {
	r := gin.Default()
	r.GET("/rooms", func(ctx *gin.Context) {
		keys := make([]RoomCreate, len(rooms))

		i := 0
		for k, v := range rooms {
			keys[i].Name = k
			keys[i].ConnectionsLength = len(v.Clients)
			i++
		}

		ctx.JSON(200, keys)
	})
	r.POST("/rooms", func(ctx *gin.Context) {
		data := &RoomCreate{}
		if err := ctx.BindJSON(data); err != nil {
			ctx.JSON(400, gin.H{
				"message": "bad body request",
			})
			return
		}

		if room, ok := rooms[data.Name]; ok {
			data.ConnectionsLength = len(room.Clients)
			ctx.JSON(200, data)
			return
		}

		room := core.NewRoom()
		go room.Run()
		rooms[data.Name] = room
		ctx.JSON(201, data)
	})
	r.DELETE("/rooms/:roomName", func(ctx *gin.Context) {
		var (
			roomName            = ctx.Param("roomName")
			room     *core.Room = nil
			ok                  = false
		)
		if room, ok = rooms[roomName]; !ok {
			ctx.Status(400)
			return
		}

		delete(rooms, roomName)
		room.Terminate()
		ctx.Status(204)
	})
	r.GET("/rooms/:roomName/ws", func(ctx *gin.Context) {
		var (
			roomName            = ctx.Param("roomName")
			room     *core.Room = nil
			ok                  = false
		)
		if room, ok = rooms[roomName]; !ok {
			ctx.Status(400)
			return
		}

		core.ConnectToRoom(room, ctx.Writer, ctx.Request)
	})
	r.Run(addr)
}
