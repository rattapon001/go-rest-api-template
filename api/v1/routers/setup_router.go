package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers"
	"github.com/rattapon001/go-rest-api-template/internal/repository"
	"github.com/rattapon001/go-rest-api-template/internal/usecase"
	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) *gin.Engine {

	batchRepo := repository.NewBatchRepository(db)
	batchUseCase := usecase.NewBatchService(batchRepo)
	batchHandler := handlers.NewBatchHandler(batchUseCase)

	allocateRepo := repository.NewAllocateRepository(db)
	allocateUseCase := usecase.NewAllocateUseCase(allocateRepo)
	allocateHandler := handlers.NewAllocateHandler(allocateUseCase)

	v1 := r.Group("api/v1")
	{
		v1.POST("/add-batch", batchHandler.AddBatch)
		v1.POST("/allocate", allocateHandler.AllocateCreate)
	}
	return r
}
