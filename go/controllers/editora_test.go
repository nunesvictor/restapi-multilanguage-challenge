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

func TestGetEditoras(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/editoras/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var editoras []models.Editora
	json.Unmarshal(body, &editoras)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(editoras), 1, fmt.Sprintf("editoras length: %d", len(editoras)))
	assert.Equal(t, editoras[0].ID, editoraTest.ID, fmt.Sprintf("editoras[0].ID: %d", editoras[0].ID))
	assert.Equal(t, editoras[0].Nome, editoraTest.Nome, fmt.Sprintf("editoras[0].Nome: %s", editoras[0].Nome))
	assert.Equal(t, editoras[0].Localidade, editoraTest.Localidade, fmt.Sprintf("editoras[0].Localidade: %s", editoras[0].Localidade))
}

func TestGetEditora(t *testing.T) {
	setup()
	defer shutdown()

	var editora models.Editora

	req, _ := http.NewRequest("GET", "/editoras/1", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &editora)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, editora.ID, editoraTest.ID, fmt.Sprintf("editora.ID: %d", editora.ID))
	assert.Equal(t, editora.Nome, editoraTest.Nome, fmt.Sprintf("editora.Nome: %s", editora.Nome))
	assert.Equal(t, editora.Localidade, editoraTest.Localidade, fmt.Sprintf("editora.Localidade: %s", editora.Localidade))
}

func TestCreateEditora(t *testing.T) {
	setup()
	defer shutdown()

	reqEditora := models.Editora{Nome: "Editora B", Localidade: "Gurupi"}
	jsonStr, _ := json.Marshal(&reqEditora)

	req, _ := http.NewRequest("POST", "/editoras/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var respEditora models.Editora
	json.Unmarshal(body, &respEditora)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, reqEditora.Nome, respEditora.Nome, fmt.Sprintf("editora.Nome: %s", reqEditora.Nome))
	assert.Equal(t, reqEditora.Localidade, respEditora.Localidade, fmt.Sprintf("editora.Localidade: %s", reqEditora.Localidade))
}

func TestUpdateEditora(t *testing.T) {
	setup()
	defer shutdown()

	reqEditora := models.Editora{Nome: "Editora C", Localidade: "Araguaína"}
	jsonStr, _ := json.Marshal(&reqEditora)

	req, _ := http.NewRequest("PUT", "/editoras/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	var respEditora models.Editora
	database.DB.First(&respEditora, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, respEditora.ID, uint(1), fmt.Sprintf("editora.ID: %d", respEditora.ID))
	assert.Equal(t, respEditora.Nome, reqEditora.Nome, fmt.Sprintf("editora.Nome: %s", respEditora.Nome))
	assert.Equal(t, respEditora.Localidade, reqEditora.Localidade, fmt.Sprintf("editora.Localidade: %s", respEditora.Localidade))
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
