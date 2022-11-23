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

func TestGetLivros(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/livros/", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var livros []models.Livro
	json.Unmarshal(body, &livros)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, len(livros), 1, fmt.Sprintf("livros length: %d", len(livros)))
	assert.Equal(t, livros[0].ID, livroTest.ID, fmt.Sprintf("livros[0].ID: %d", livros[0].ID))
	assert.Equal(t, livros[0].Titulo, livroTest.Titulo, fmt.Sprintf("livros[0].Titulo: %s", livros[0].Titulo))
	assert.Equal(t, livros[0].EditoraID, livroTest.EditoraID, fmt.Sprintf("livros[0].EditoraID: %d", livros[0].EditoraID))
	assert.Equal(t, livros[0].Editora.Nome, livroTest.Editora.Nome, fmt.Sprintf("livros[0].Editora.Nome: %s", livros[0].Editora.Nome))
	assert.Equal(t, livros[0].Editora.Localidade, livroTest.Editora.Localidade, fmt.Sprintf("livros[0].Editora.Localidade: %s", livros[0].Editora.Localidade))
	assert.Equal(t, livros[0].Autores[0].ID, livroTest.Autores[0].ID, fmt.Sprintf("livros[0].Autores[0].ID: %d", livros[0].Autores[0].ID))
	assert.Equal(t, livros[0].Autores[0].Nome, livroTest.Autores[0].Nome, fmt.Sprintf("livros[0].Autores[0].Nome: %s", livros[0].Autores[0].Nome))
	assert.Equal(t, livros[0].Autores[0].Sobrenome, livroTest.Autores[0].Sobrenome, fmt.Sprintf("livros[0].Autores[0].Sobrenome: %s", livros[0].Autores[0].Sobrenome))
	assert.Equal(t, livros[0].Generos[0].ID, livroTest.Generos[0].ID, fmt.Sprintf("livros[0].Generos[0].ID: %d", livros[0].Generos[0].ID))
	assert.Equal(t, livros[0].Generos[0].Descricao, livroTest.Generos[0].Descricao, fmt.Sprintf("livros[0].Generos[0].Descricao: %s", livros[0].Generos[0].Descricao))
	assert.Equal(t, livros[0].Edicao, livroTest.Edicao, fmt.Sprintf("livros[0].Edicao: %d", livros[0].Edicao))
	assert.Equal(t, livros[0].Ano, livroTest.Ano, fmt.Sprintf("livros[0].Ano: %d", livros[0].Ano))
	assert.Equal(t, livros[0].Paginas, livroTest.Paginas, fmt.Sprintf("livros[0].Paginas: %d", livros[0].Paginas))
	assert.Equal(t, livros[0].ISBN, livroTest.ISBN, fmt.Sprintf("livros[0].ISBN: %s", livros[0].ISBN))
}

func TestGetLivro(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("GET", "/livros/1", nil)
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var livro models.Livro
	json.Unmarshal(body, &livro)

	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, livro.ID, livroTest.ID, fmt.Sprintf("livro.ID: %d", livro.ID))
	assert.Equal(t, livro.Titulo, livroTest.Titulo, fmt.Sprintf("livro.Titulo: %s", livro.Titulo))
	assert.Equal(t, livro.EditoraID, livroTest.EditoraID, fmt.Sprintf("livro.EditoraID: %d", livro.EditoraID))
	assert.Equal(t, livro.Editora.Nome, livroTest.Editora.Nome, fmt.Sprintf("livro.Editora.Nome: %s", livro.Editora.Nome))
	assert.Equal(t, livro.Editora.Localidade, livroTest.Editora.Localidade, fmt.Sprintf("livro.Editora.Localidade: %s", livro.Editora.Localidade))
	assert.Equal(t, livro.Autores[0].ID, livroTest.Autores[0].ID, fmt.Sprintf("livro.Autores[0].ID: %d", livro.Autores[0].ID))
	assert.Equal(t, livro.Autores[0].Nome, livroTest.Autores[0].Nome, fmt.Sprintf("livro.Autores[0].Nome: %s", livro.Autores[0].Nome))
	assert.Equal(t, livro.Autores[0].Sobrenome, livroTest.Autores[0].Sobrenome, fmt.Sprintf("livro.Autores[0].Sobrenome: %s", livro.Autores[0].Sobrenome))
	assert.Equal(t, livro.Generos[0].ID, livroTest.Generos[0].ID, fmt.Sprintf("livro.Generos[0].ID: %d", livro.Generos[0].ID))
	assert.Equal(t, livro.Generos[0].Descricao, livroTest.Generos[0].Descricao, fmt.Sprintf("livro.Generos[0].Descricao: %s", livro.Generos[0].Descricao))
	assert.Equal(t, livro.Edicao, livroTest.Edicao, fmt.Sprintf("livro.Edicao: %d", livro.Edicao))
	assert.Equal(t, livro.Ano, livroTest.Ano, fmt.Sprintf("livro.Ano: %d", livro.Ano))
	assert.Equal(t, livro.Paginas, livroTest.Paginas, fmt.Sprintf("livro.Paginas: %d", livro.Paginas))
	assert.Equal(t, livro.ISBN, livroTest.ISBN, fmt.Sprintf("livro.ISBN: %s", livro.ISBN))
}

func TestCreateLivro(t *testing.T) {
	setup()
	defer shutdown()

	reqLivro := models.Livro{
		Titulo:    "Livro B",
		EditoraID: 1,
		Autores: []models.Autor{{
			ID: 1,
		}},
		Generos: []models.Genero{{
			ID: 1,
		}},
		Edicao:  2,
		Ano:     2020,
		Paginas: 200,
		ISBN:    "1234567890",
	}
	jsonStr, _ := json.Marshal(&reqLivro)

	req, _ := http.NewRequest("POST", "/livros/", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	var respLivro models.Livro
	json.Unmarshal(body, &respLivro)

	assert.Equal(t, w.Code, http.StatusCreated, w.Body.String())
	assert.Equal(t, respLivro.Titulo, reqLivro.Titulo, fmt.Sprintf("livro.Titulo: %s", respLivro.Titulo))
	assert.Equal(t, respLivro.EditoraID, reqLivro.EditoraID, fmt.Sprintf("livro.EditoraID: %d", respLivro.EditoraID))
	assert.Equal(t, respLivro.Editora.Nome, reqLivro.Editora.Nome, fmt.Sprintf("livro.Editora.Nome: %s", respLivro.Editora.Nome))
	assert.Equal(t, respLivro.Editora.Localidade, reqLivro.Editora.Localidade, fmt.Sprintf("livro.Editora.Localidade: %s", respLivro.Editora.Localidade))
	assert.Equal(t, respLivro.Autores[0].ID, reqLivro.Autores[0].ID, fmt.Sprintf("livro.Autores[0].ID: %d", respLivro.Autores[0].ID))
	assert.Equal(t, respLivro.Generos[0].ID, reqLivro.Generos[0].ID, fmt.Sprintf("livro.Generos[0].ID: %d", respLivro.Generos[0].ID))
	assert.Equal(t, respLivro.Edicao, reqLivro.Edicao, fmt.Sprintf("livro.Edicao: %d", respLivro.Edicao))
	assert.Equal(t, respLivro.Ano, reqLivro.Ano, fmt.Sprintf("livro.Ano: %d", respLivro.Ano))
	assert.Equal(t, respLivro.Paginas, reqLivro.Paginas, fmt.Sprintf("livro.Paginas: %d", respLivro.Paginas))
	assert.Equal(t, respLivro.ISBN, reqLivro.ISBN, fmt.Sprintf("livro.ISBN: %s", respLivro.ISBN))
}

func TestUpdateLivro(t *testing.T) {
	setup()
	defer shutdown()

	autor := models.Autor{Nome: "José", Sobrenome: "Maria Martins"}
	database.DB.Create(&autor)

	genero := models.Genero{Descricao: "Romance"}
	database.DB.Create(&genero)

	reqLivro := models.Livro{
		Titulo:    "Livro C",
		EditoraID: 1,
		Autores:   []models.Autor{autor},
		Generos:   []models.Genero{genero},
		Edicao:    4,
		Ano:       2019,
		Paginas:   250,
		ISBN:      "A1B2C3D4F5G6",
	}
	jsonStr, _ := json.Marshal(&reqLivro)

	req, _ := http.NewRequest("PUT", "/livros/1", bytes.NewBuffer(jsonStr))
	r.ServeHTTP(w, req)

	var respLivro models.Livro
	database.DB.
		Preload("Autores").
		Preload("Generos").
		First(&respLivro, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, respLivro.ID, uint(1), fmt.Sprintf("livro.ID: %d", respLivro.ID))
	assert.Equal(t, respLivro.Titulo, reqLivro.Titulo, fmt.Sprintf("livro.Titulo: %s", respLivro.Titulo))
	assert.Equal(t, respLivro.EditoraID, reqLivro.EditoraID, fmt.Sprintf("livro.EditoraID: %d", respLivro.EditoraID))
	assert.Equal(t, respLivro.Editora.Nome, reqLivro.Editora.Nome, fmt.Sprintf("livro.Editora.Nome: %s", respLivro.Editora.Nome))
	assert.Equal(t, respLivro.Editora.Localidade, reqLivro.Editora.Localidade, fmt.Sprintf("livro.Editora.Localidade: %s", respLivro.Editora.Localidade))
	assert.Equal(t, respLivro.Autores[0].ID, reqLivro.Autores[0].ID, fmt.Sprintf("livro.Autores[0].ID: %d", respLivro.Autores[0].ID))
	assert.Equal(t, respLivro.Generos[0].ID, reqLivro.Generos[0].ID, fmt.Sprintf("livro.Generos[0].ID: %d", respLivro.Generos[0].ID))
	assert.Equal(t, respLivro.Edicao, reqLivro.Edicao, fmt.Sprintf("livro.Edicao: %d", respLivro.Edicao))
	assert.Equal(t, respLivro.Ano, reqLivro.Ano, fmt.Sprintf("livro.Ano: %d", respLivro.Ano))
	assert.Equal(t, respLivro.Paginas, reqLivro.Paginas, fmt.Sprintf("livro.Paginas: %d", respLivro.Paginas))
	assert.Equal(t, respLivro.ISBN, reqLivro.ISBN, fmt.Sprintf("livro.ISBN: %s", respLivro.ISBN))
}

func TestDeleteLivro(t *testing.T) {
	setup()
	defer shutdown()

	req, _ := http.NewRequest("DELETE", "/livros/1", nil)
	r.ServeHTTP(w, req)

	var livro models.Livro
	database.DB.First(&livro, uint(1))

	assert.Equal(t, w.Code, http.StatusNoContent)
	assert.Equal(t, livro.ID, uint(0))
}
