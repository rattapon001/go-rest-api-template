package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
	"github.com/rattapon001/go-rest-api-template/internal/usecase"
)

type allocateHandler struct {
	allocateUseCase usecase.AllocateUseCase
}

func NewAllocateHandler(usecase usecase.AllocateUseCase) *allocateHandler {
	return &allocateHandler{allocateUseCase: usecase}
}

func (h *allocateHandler) AllocateCreate(c *gin.Context) {
	allocateInput := dto.AllocateInput{}
	if err := c.ShouldBindJSON(&allocateInput); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	h.allocateUseCase.AllocateCreate(&allocateInput)
	c.JSON(200, nil)
}
