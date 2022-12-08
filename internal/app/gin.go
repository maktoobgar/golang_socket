package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/configs"
	"github.com/maktoobgar/golang_socket/internal/routers"
)

func Gin() {
	r := gin.Default()

	routers.ApplyAllRouters(r)

	r.Run(fmt.Sprintf("%s:%d", configs.CFG.Host, configs.CFG.Port))
}
