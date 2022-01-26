package router

import (
	"github.com/RalapZ/DeepBluePaas/common/middleware"
	"github.com/RalapZ/DeepBluePaas/server/config"
	"github.com/gin-gonic/gin"
)

type DeepBlueRouterGroup struct{
	DP_CASBIN
	DP_AUTH
}





func RouterConfig(e *gin.Engine){
	e.Use(middleware.GinLogger())
	//Server.Use(middleware.PringTest())
	e.Use(middleware.JWTAuth(config.GVA_REDIS,config.GVA_DB,&config.LocalBlackCache))


	//Server.Use()
	e.GET("/hello",hello)
	Private:= e.Group("/api")
	{

	}

	group := e.Group("/api/v1/auth")
	{
		group.POST("/user",UserAdd)
		group.GET("/user/:id",UserGetDetail)
		group.GET("/user/list",UserGetList)
		group.DELETE("/user/:id",UserDelete)
	}
}