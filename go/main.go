package main

import (
	"gin-restapi/database"
	"gin-restapi/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupEnv(filenames ...string) {
	if err := godotenv.Load(filenames...); err != nil {
		cwd, _ := os.Getwd()
		log.Fatalf("%s/.env: arquivo não encontrado", cwd)
	}
}

func main() {
	setupEnv()

	r := router.SetupRouter(database.Connect)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run()
}
