package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-restapi/database"
	"gin-restapi/models"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGeneros(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/generos/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var generos []models.Genero
	json.Unmarshal(body, &generos)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(generos), 1, fmt.Sprintf("generos length: %d", len(generos)))
	assert.Equal(t, generos[0].ID, generoTest.ID, fmt.Sprintf("generos[0].ID: %d", generos[0].ID))
	assert.Equal(t, generos[0].Descricao, generoTest.Descricao, fmt.Sprintf("generos[0].Descricao: %s", generos[0].Descricao))
}

func TestGetGenero(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/generos/1", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var genero models.Genero
	json.Unmarshal(body, &genero)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, genero.ID, generoTest.ID, fmt.Sprintf("genero.ID: %d", genero.ID))
	assert.Equal(t, genero.Descricao, generoTest.Descricao, fmt.Sprintf("genero.Descricao: %s", genero.Descricao))
}

func TestCreateGenero(t *testing.T) {
	setup()
	defer shutdown()

	reqGenero := models.Genero{Descricao: "Genero B"}
	jsonStr, _ := json.Marshal(&reqGenero)

	req, _ := http.NewRequest("POST", "/generos/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var respGenero models.Genero
	json.Unmarshal(body, &respGenero)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, respGenero.Descricao, reqGenero.Descricao, fmt.Sprintf("genero.Descricao: %s", respGenero.Descricao))
}

func TestUpdateGenero(t *testing.T) {
	setup()
	defer shutdown()

	reqGenero := models.Genero{Descricao: "Genero B"}
	jsonStr, _ := json.Marshal(&reqGenero)

	req, _ := http.NewRequest("PUT", "/generos/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	var respGenero models.Genero
	database.DB.First(&respGenero, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, respGenero.ID, uint(1), fmt.Sprintf("genero.ID: %d", respGenero.ID))
	assert.Equal(t, respGenero.Descricao, reqGenero.Descricao, fmt.Sprintf("genero.Descricao: %s", respGenero.Descricao))
}

func TestDeleteGenero(t *testing.T) {
	setup()
	defer shutdown()

	genero := models.Genero{Descricao: "Genero B"}
	jsonStr, _ := json.Marshal(&genero)

	req, _ := http.NewRequest("DELETE", "/generos/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	database.DB.First(&models.Genero{}, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, genero.ID, uint(0))
}
