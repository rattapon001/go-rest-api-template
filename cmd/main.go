package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rattapon001/go-rest-api-template/api/v1/routers"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	port := os.Getenv("PORT")
	r := gin.Default()
	r.GET("/ping")
	routers.Setup(r)
	r.Run(":" + port)
}
