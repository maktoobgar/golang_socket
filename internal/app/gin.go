package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/maktoobgar/golang_socket/internal/configs"
	"github.com/maktoobgar/golang_socket/internal/routers"
)

func configCors(r *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowWebSockets = true
	if configs.CFG.Debug {
		corsConfig.AllowOriginFunc = func(origin string) bool {
			return true
		}
	} else {
		corsConfig.AllowOrigins = configs.CFG.AllowOrigins
	}

	r.Use(cors.New(corsConfig))
}

func Gin() {
	if !configs.CFG.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	configCors(r)

	routers.ApplyAllRouters(r)

	r.Run(fmt.Sprintf("%s:%d", configs.CFG.Host, configs.CFG.Port))
}
