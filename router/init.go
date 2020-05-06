package router

import (
	"github.com/gin-gonic/gin"
	"lijinbo/middleware"
	"lijinbo/model"
	"net/http"
)


func Init(ge *gin.Engine)  {
	recordGroup := ge.Group("/record").Use(middleware.UserLogin())
	{
		recordGroup.GET("/getOne", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data":model.Record{},
			})
		})
		recordGroup.GET("/lists", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data":map[int] model.Record {
					0:model.Record{},
				},
			})
		})
	}

	userGroup := ge.Group("/user")
	{
		userGroup.GET("/records", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": model.Record{},
			})
		})
	}
}
