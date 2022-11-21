package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"gin-restapi/database"
	"gin-restapi/models"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetEditoras(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/editoras/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), "Editora A")
	assert.Contains(t, w.Body.String(), "Palmas")
}

func TestGetEditora(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/editoras/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), "Editora A")
	assert.Contains(t, w.Body.String(), "Palmas")
}

func TestCreateEditora(t *testing.T) {
	setup()
	defer shutdown()

	editora := models.Editora{Nome: "Editora B", Localidade: "Gurupi"}
	jsonStr, _ := json.Marshal(&editora)

	req, _ := http.NewRequest("POST", "/editoras/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &editora)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, editora.ID, uint(2))
	assert.Equal(t, editora.Nome, "Editora B")
	assert.Equal(t, editora.Localidade, "Gurupi")
}

func TestUpdateEditora(t *testing.T) {
	setup()
	defer shutdown()

	autor := models.Editora{Nome: "Editora C", Localidade: "Araguaína"}
	jsonStr, _ := json.Marshal(&autor)

	req, _ := http.NewRequest("PUT", "/editoras/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	database.DB.First(&autor, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, autor.Nome, "Editora C")
	assert.Equal(t, autor.Localidade, "Araguaína")
}

func TestDeleteEditora(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("DELETE", "/editoras/1", nil)
	r.ServeHTTP(w, req)

	result := database.DB.First(&models.Editora{}, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, errors.Is(result.Error, gorm.ErrRecordNotFound), true)
}
