package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"gin-restapi/database"
	"gin-restapi/models"
	"gin-restapi/router"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
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
}

func shutdown() {
	if err := os.Remove("test.db"); err != nil {
		log.Panic("Erro ao excluir arquivo de banco de dados")
	}
}

func TestGetAutor(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/autores/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetAutores(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/autores/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), "Fulano")
	assert.Contains(t, w.Body.String(), "de Tal")
}

func TestCreateAutor(t *testing.T) {
	setup()
	defer shutdown()

	autor := models.Autor{Nome: "Beltrano", Sobrenome: "de Tal"}
	jsonStr, _ := json.Marshal(&autor)

	req, _ := http.NewRequest("POST", "/autores/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &autor)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, autor.ID, uint(2))
	assert.Equal(t, autor.Nome, "Beltrano")
	assert.Equal(t, autor.Sobrenome, "de Tal")
}

func TestUpdateAutor(t *testing.T) {
	setup()
	defer shutdown()

	autor := models.Autor{Nome: "Ciclano", Sobrenome: "da Silva"}
	jsonStr, _ := json.Marshal(&autor)

	req, _ := http.NewRequest("PUT", "/autores/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	database.DB.First(&autor, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, autor.Nome, "Ciclano")
	assert.Equal(t, autor.Sobrenome, "da Silva")
}

func TestDeleteAutor(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("DELETE", "/autores/1", nil)
	r.ServeHTTP(w, req)

	result := database.DB.First(&models.Autor{}, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, errors.Is(result.Error, gorm.ErrRecordNotFound), true)
}
