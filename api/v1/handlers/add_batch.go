package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
	"github.com/rattapon001/go-rest-api-template/internal/usecase"
)

type batchHandler struct {
	batchUseCase usecase.BatchUseCase
}

func NewBatchHandler(usecase usecase.BatchUseCase) *batchHandler {
	return &batchHandler{batchUseCase: usecase}
}

func (h batchHandler) AddBatch(c *gin.Context) {

	batch := dto.AddBatchInput{}
	if err := c.ShouldBindJSON(&batch); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	h.batchUseCase.AddBatch(&batch)
	log.Println(batch)
	c.JSON(200, nil)
}
