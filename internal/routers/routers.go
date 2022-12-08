package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/routers/http"
	"github.com/maktoobgar/golang_socket/internal/routers/ws"
)

func ApplyAllRouters(r *gin.Engine) {
	http.ApplyHttpRouters(r)
	ws.ApplyWsRouters(r)
}
