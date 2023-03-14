package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers/dto"
	"github.com/rattapon001/go-rest-api-template/internal/usecase"
	"github.com/rattapon001/go-rest-api-template/pkg/enum"
	"github.com/rattapon001/go-rest-api-template/pkg/helpers"
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
		helpers.Response(c, http.StatusBadRequest, string(enum.FIAL), []interface{}{})
		return
	}
	allocate, err := h.allocateUseCase.AllocateCreate(&allocateInput)
	if err != nil {
		helpers.Response(c, http.StatusBadRequest, string(enum.FIAL), []interface{}{})
	} else {
		helpers.Response(c, http.StatusOK, string(enum.SUCCESS), []interface{}{allocate})
	}
}
