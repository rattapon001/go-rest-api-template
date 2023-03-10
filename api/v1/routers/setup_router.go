package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rattapon001/go-rest-api-template/api/v1/handlers"
	"github.com/rattapon001/go-rest-api-template/internal/db"
	"github.com/rattapon001/go-rest-api-template/internal/repository"
	"github.com/rattapon001/go-rest-api-template/internal/usecase"
	"gorm.io/driver/postgres"
)

func Setup() *gin.Engine {

	dsn := os.Getenv("DATA_BASE")

	dbSession := db.ConnectDatabase(postgres.Open(dsn))

	batchRepo := repository.NewBatchRepository(dbSession)
	batchUseCase := usecase.NewServiceCreate(batchRepo)
	batchHandler := handlers.NewBatchHandler(batchUseCase)

	r := gin.Default()
	r.GET("/ping")

	v1 := r.Group("api/v1")
	{
		v1.POST("/add-batch", batchHandler.AddBatch)
		v1.POST("/allocate", handlers.Allocate)
	}
	return r
}
