package http

import "github.com/gin-gonic/gin"

func ApplyHttpRouters(r *gin.Engine) {
	roomsResource(r)
}
