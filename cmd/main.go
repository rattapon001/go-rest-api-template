package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rattapon001/go-rest-api-template/api/v1/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	dsn := os.Getenv("DATA_BASE")

	dbSession, err := gorm.Open(postgres.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})

	if err != nil {
		log.Println("Connect into Database Failed")
	}

	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/ping")
	routers.Setup(r, dbSession)
	r.Run(":" + port)
}
