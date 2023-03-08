package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rattapon001/go-rest-api-template/api/v1/routers"
	// "github.com/rattapon001/go-rest-api-template/internal/db"
	// "gorm.io/driver/postgres"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	port := os.Getenv("PORT")
	// dsn := os.Getenv("DATA_BASE")

	// dbSession := db.ConnectDatabase(postgres.Open(dsn))

	r := routers.Setup()
	r.Run(":" + port)
}
