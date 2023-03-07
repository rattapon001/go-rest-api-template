package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
)

func AddBacth(c *gin.Context) {
	bacth := dto.AddBacthInput{}
	if err := c.ShouldBindJSON(&bacth); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	log.Println(bacth)

	c.JSON(200, nil)
}
