package main

import (
	"gin-restapi/src/database"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		cwd, _ := os.Getwd()
		log.Fatalf("%s/.env: arquivo não encontrad'o", cwd)
	}
}

func main() {
	loadEnv()

	database.Connect()
	r := gin.Default()

	log.Printf("Listening on http://localhost:8000\n\n")
	http.ListenAndServe(":8000", r)
}
