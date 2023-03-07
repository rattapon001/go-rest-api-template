package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
)

func AddBatch(c *gin.Context) {
	batch := dto.AddBatchInput{}
	if err := c.ShouldBindJSON(&batch); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	log.Println(batch)

	c.JSON(200, nil)
}
