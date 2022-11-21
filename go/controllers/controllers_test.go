package controllers_test

import (
	"gin-restapi/database"
	"gin-restapi/models"
	"gin-restapi/router"
	"log"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
	w *httptest.ResponseRecorder
)

func setup() {
	gin.SetMode(gin.ReleaseMode)

	r = router.SetupRouter(database.ConnectTest)
	w = httptest.NewRecorder()

	database.DB.Create(&models.Autor{Nome: "Fulano", Sobrenome: "de Tal"})
	database.DB.Create(&models.Editora{Nome: "Editora A", Localidade: "Palmas"})
	database.DB.Create(&models.Genero{Descricao: "Genero A"})
}

func shutdown() {
	if err := os.Remove("test.db"); err != nil {
		log.Panic("Erro ao excluir arquivo de banco de dados")
	}
}
