package database

import (
	"gin-restapi/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DATABASE_DSN")), &gorm.Config{})

	if err != nil {
		os.Exit(1)
	}

	DB.AutoMigrate(&models.Autor{}, &models.Editora{}, &models.Genero{}, &models.Livro{})
}
