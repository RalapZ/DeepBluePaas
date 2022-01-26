package router

import (
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/RalapZ/DeepBluePaas/server/config"
	"github.com/gin-gonic/gin"
)




func Serve(){
	err := middleware.InitLogger(&config.LogConfig)
	config.RedisGetClientConn()
	if err != nil{
		panic(err)
	}
	defer middleware.LogGlobal.Sync()
	Server := gin.Default()

	RouterConfig(Server)

	Server.Run("127.0.0.1:18080")
}
