package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.GET("/ping")

	v1 := r.Group("api/v1")
	{
		v1.POST("/add-bacth", handlers.AddBacth)
		v1.POST("/allocate", handlers.Allocate)
	}
	return r
}
