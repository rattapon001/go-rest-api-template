package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rattapon001/go-rest-api-template/api/v1/routers"
)

func main() {
	err := godotenv.Load("./configs/local.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	port := os.Getenv("PORT")

	r := routers.Setup()
	r.Run(":" + port)
}
