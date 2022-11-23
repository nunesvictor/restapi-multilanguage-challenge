package controllers

import (
	"gin-restapi/database"
	"gin-restapi/httputil"
	"gin-restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /
//
// @Summary Lista livros
// @Schemes
// @Description Recupera a lista de livros
// @Tags livros
// @Accept json
// @Produce json
// @Success 200 {array} models.Livro
// @Failure 500 {object} httputil.HTTPError
// @Router /livros/ [get]
func GetLivros(c *gin.Context) {
	var livros []models.Livro
	db := database.DB.
		Preload("Editora").
		Preload("Autores").
		Preload("Generos")

	models.ListAll(c, db, &livros)
	c.JSON(http.StatusOK, &livros)
}

// @BasePath /
//
// @Summary Recupera livro
// @Schemes
// @Description Recupera um livro por id
// @Tags livros
// @Accept json
// @Produce json
// @Param id path int true "ID do livro"
// @Success 200 {object} models.Livro
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /livros/{id} [get]
func GetLivro(c *gin.Context) {
	var livro models.Livro
	db := database.DB.
		Preload("Editora").
		Preload("Autores").
		Preload("Generos")

	models.Get(c, db, &livro)
	c.JSON(http.StatusOK, &livro)
}

// @BasePath /
//
// @Summary Cria livro
// @Schemes
// @Description Cria um novo livro
// @Tags livros
// @Accept json
// @Produce json
// @Param livro body models.Livro true "Dados do livro"
// @Success 201 {object} models.Livro
// @Failure 500 {object} httputil.HTTPError
// @Router /livros/ [post]
func CreateLivro(c *gin.Context) {
	var livro models.Livro
	db := database.DB.Preload("Autores").Preload("Generos")

	if err := c.ShouldBindJSON(&livro); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	database.DB.Find(&livro.Autores, &livro.Autores[0].ID)
	database.DB.Find(&livro.Generos, &livro.Generos[0].ID)

	models.CreateLivro(c, db, &livro)

	c.JSON(http.StatusCreated, &livro)
}

// @BasePath /
//
// @Summary Atualiza livro
// @Schemes
// @Description Atualiza os dados de um livro
// @Tags livros
// @Accept json
// @Produce json
// @Param id path int true "ID do livro"
// @Param livro body models.Livro true "Dados do livro"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /livros/{id} [put]
func UpdateLivro(c *gin.Context) {
	var livro, livroInput models.Livro
	models.Get(c, database.DB, &livro)

	if err := c.ShouldBindJSON(&livroInput); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	database.DB.Find(&livroInput.Autores, &livroInput.Autores[0].ID)
	database.DB.Find(&livroInput.Generos, &livroInput.Generos[0].ID)

	livro.Update(c, database.DB, &livroInput)
	c.JSON(http.StatusNoContent, nil)
}

// @BasePath /
//
// @Summary Deleta livro
// @Schemes
// @Description Remove o cadastro de um livro
// @Tags livros
// @Accept json
// @Produce json
// @Param id path int true "ID do livro"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /livros/{id} [delete]
func DeleteLivro(c *gin.Context) {
	models.Delete[models.Livro](c, database.DB)
	c.JSON(http.StatusNoContent, nil)
}
