package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
	"github.com/rattapon001/go-rest-api-template/internal/usecase"
	"github.com/rattapon001/go-rest-api-template/pkg/enum"
	"github.com/rattapon001/go-rest-api-template/pkg/helpers"
)

type batchHandler struct {
	batchUseCase usecase.BatchUseCase
}

func NewBatchHandler(usecase usecase.BatchUseCase) *batchHandler {
	return &batchHandler{batchUseCase: usecase}
}

func (h *batchHandler) AddBatch(c *gin.Context) {

	batchInput := dto.AddBatchInput{}
	if err := c.ShouldBindJSON(&batchInput); err != nil {
		helpers.Response(c, http.StatusBadRequest, string(enum.FIAL), []interface{}{})
		return
	}
	batch, err := h.batchUseCase.AddBatch(&batchInput)
	if err != nil {
		helpers.Response(c, http.StatusBadRequest, string(enum.FIAL), []interface{}{})
	} else {
		helpers.Response(c, http.StatusOK, string(enum.SUCCESS), []interface{}{batch})
	}

}
