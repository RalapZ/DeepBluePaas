package router

import (
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context){
	c.String(200,"hello")
}