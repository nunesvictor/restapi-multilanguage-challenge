package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gin-restapi/database"
	"gin-restapi/models"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetAutores(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/autores/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var autores []models.Autor
	json.Unmarshal(body, &autores)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(autores), 1, fmt.Sprintf("autores length: %d", len(autores)))
	assert.Equal(t, autores[0].ID, autorTest.ID, fmt.Sprintf("autores[0].ID: %d", autores[0].ID))
	assert.Equal(t, autores[0].Nome, autorTest.Nome, fmt.Sprintf("autores[0].Nome: %s", autores[0].Nome))
	assert.Equal(t, autores[0].Sobrenome, autorTest.Sobrenome, fmt.Sprintf("autores[0].Sobrenome: %s", autores[0].Sobrenome))
}

func TestGetAutor(t *testing.T) {
	setup()
	defer shutdown()

	var autor models.Autor

	req, _ := http.NewRequest("GET", "/autores/1", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &autor)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, autor.ID, autorTest.ID, fmt.Sprintf("autor.ID: %d", autor.ID))
	assert.Equal(t, autor.Nome, autorTest.Nome, fmt.Sprintf("autor.Nome: %s", autor.Nome))
	assert.Equal(t, autor.Sobrenome, autorTest.Sobrenome, fmt.Sprintf("autor.Sobrenome: %s", autor.Sobrenome))
}

func TestCreateAutor(t *testing.T) {
	setup()
	defer shutdown()

	reqAutor := models.Autor{Nome: "Beltrano", Sobrenome: "de Tal"}
	jsonStr, _ := json.Marshal(&reqAutor)

	req, _ := http.NewRequest("POST", "/autores/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var respAutor models.Autor
	json.Unmarshal(body, &respAutor)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, respAutor.Nome, reqAutor.Nome, fmt.Sprintf("autor.Nome: %s", respAutor.Nome))
	assert.Equal(t, respAutor.Sobrenome, reqAutor.Sobrenome, fmt.Sprintf("autor.Sobrenome: %s", respAutor.Sobrenome))
}

func TestUpdateAutor(t *testing.T) {
	setup()
	defer shutdown()

	reqAutor := models.Autor{Nome: "Ciclano", Sobrenome: "da Silva"}
	jsonStr, _ := json.Marshal(&reqAutor)

	req, _ := http.NewRequest("PUT", "/autores/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	var respAutor models.Autor
	database.DB.First(&respAutor, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, respAutor.ID, uint(1), fmt.Sprintf("autor.ID: %d", respAutor.ID))
	assert.Equal(t, respAutor.Nome, reqAutor.Nome, fmt.Sprintf("autor.Nome: %s", respAutor.Nome))
	assert.Equal(t, respAutor.Sobrenome, reqAutor.Sobrenome, fmt.Sprintf("autor.Sobrenome: %s", respAutor.Sobrenome))
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
