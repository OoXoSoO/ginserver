package main

import (
	"ginserver/pkg"
	"ginserver/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ps := &usecase.PongService{}
	cs := &usecase.CreateService{}
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": ps.Pong(c),
		})
	})
	engine.POST("/user", func(c *gin.Context) {
		rq := pkg.UserCreateInput{}
		c.Bind(&rq)
		rs, err := cs.Create(c, rq)
		if err != nil {
			c.AbortWithError(400, err)
		}
		c.JSON(200, rs)
	})
	engine.Run() // listen and serve on 0.0.0.0:8080
}
