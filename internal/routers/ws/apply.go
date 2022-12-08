package ws

import "github.com/gin-gonic/gin"

func ApplyWsRouters(r *gin.Engine) {
	joinRooms(r)
}
