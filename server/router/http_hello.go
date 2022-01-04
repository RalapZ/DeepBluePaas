package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context){
	fmt.Printf("%##v",c)
	c.String(200,"hello")

}