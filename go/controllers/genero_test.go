package controllers_test

import (
	"bytes"
	"encoding/json"
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

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), "Genero A")
}

func TestGetGenero(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/generos/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Contains(t, w.Body.String(), "Genero A")
}

func TestCreateGenero(t *testing.T) {
	setup()
	defer shutdown()

	genero := models.Genero{Descricao: "Genero B"}
	jsonStr, _ := json.Marshal(&genero)

	req, _ := http.NewRequest("POST", "/generos/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &genero)

	assert.Equal(t, w.Code, http.StatusCreated)
	assert.Equal(t, genero.ID, uint(2))
	assert.Equal(t, genero.Descricao, "Genero B")
}

func TestUpdateGenero(t *testing.T) {
	setup()
	defer shutdown()

	genero := models.Genero{Descricao: "Genero B"}
	jsonStr, _ := json.Marshal(&genero)

	req, _ := http.NewRequest("PUT", "/generos/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	database.DB.First(&models.Genero{}, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, genero.Descricao, "Genero B")
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
