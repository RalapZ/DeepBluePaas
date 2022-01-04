package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func PringTest() gin.HandlerFunc{
	return func(c *gin.Context){
		n:=time.Now()
		timedur:=time.Since(n)
		fmt.Printf("time cost is %v\n",timedur)
		c.Next()
	}
}
