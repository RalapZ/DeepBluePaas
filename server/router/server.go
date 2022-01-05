package router

import (
	"fmt"
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/gin-gonic/gin"
)




func Serve(){
	err := middleware.InitLogger()
	if err != nil{
		panic(err)
	}
	fmt.Printf("main:----%##v\n",middleware.LogGlobal)
	defer middleware.LogGlobal.Sync()
	Server := gin.New()
	//Logger, err := zap.NewProduction()
	//if err != nil{
	//	fmt.Println(err)
	//}
	//gin.Default()
	fmt.Printf("server:---%##v",middleware.LogGlobal)
	Server.Use(middleware.GinLogger())
	Server.Use(middleware.PringTest())

	//Server.Use()
	Server.GET("/hello",hello)
	Server.Run("127.0.0.1:18080")
}
