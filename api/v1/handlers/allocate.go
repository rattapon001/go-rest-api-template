package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
)

func Allocate(c *gin.Context) {

	allocate := dto.AllocateInput{}
	if err := c.ShouldBindJSON(&allocate); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	log.Println(allocate)

	c.JSON(200, nil)
}
