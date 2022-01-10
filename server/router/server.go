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
	//fmt.Printf("main:----%##v\n",middleware.LogGlobal)
	defer middleware.LogGlobal.Sync()
	Server := gin.New()
	//Logger, err := zap.NewProduction()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//gin.Default()
	Server.Use(middleware.GinLogger())
	Server.Use(middleware.PringTest())
	Server.Use(middleware.JWTAuth(config.GVA_REDIS,config.GVA_DB,&config.LocalBlackCache))


	//Server.Use()
	Server.GET("/hello",hello)
	Server.Run("127.0.0.1:18080")
}
