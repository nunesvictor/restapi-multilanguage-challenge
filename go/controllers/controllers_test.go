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
	r           *gin.Engine
	w           *httptest.ResponseRecorder
	autorTest   *models.Autor
	editoraTest *models.Editora
	generoTest  *models.Genero
)

func setup() {
	gin.SetMode(gin.ReleaseMode)

	r = router.SetupRouter(database.ConnectTest)
	w = httptest.NewRecorder()

	autorTest = &models.Autor{Nome: "Fulano", Sobrenome: "de Tal"}
	editoraTest = &models.Editora{Nome: "Editora A", Localidade: "Palmas"}
	generoTest = &models.Genero{Descricao: "Genero A"}

	database.DB.Create(autorTest)
	database.DB.Create(editoraTest)
	database.DB.Create(generoTest)
}

func shutdown() {
	if err := os.Remove("test.db"); err != nil {
		log.Panic("Erro ao excluir arquivo de banco de dados")
	}
}
