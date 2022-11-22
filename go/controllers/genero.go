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
// @Summary Lista generos
// @Schemes
// @Description Recupera a lista de generos
// @Tags generos
// @Accept json
// @Produce json
// @Success 200 {array} models.Genero
// @Failure 500 {object} httputil.HTTPError
// @Router /generos/ [get]
func GetGeneros(c *gin.Context) {
	var generos []models.Genero

	models.ListAll(c, database.DB, &generos)
	c.JSON(http.StatusOK, &generos)
}

// @BasePath /
//
// @Summary Recupera genero
// @Schemes
// @Description Recupera um genero por id
// @Tags generos
// @Accept json
// @Produce json
// @Param id path int true "ID do genero"
// @Success 200 {object} models.Genero
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /generos/{id} [get]
func GetGenero(c *gin.Context) {
	var genero models.Genero

	models.Get(c, database.DB, &genero)
	c.JSON(http.StatusOK, &genero)
}

// @BasePath /
//
// @Summary Cria genero
// @Schemes
// @Description Cria um novo genero
// @Tags generos
// @Accept json
// @Produce json
// @Param genero body models.Genero true "Dados do genero"
// @Success 201 {object} models.Genero
// @Failure 500 {object} httputil.HTTPError
// @Router /generos/ [post]
func CreateGenero(c *gin.Context) {
	var genero models.Genero

	models.Create(c, database.DB, &genero)
	c.JSON(http.StatusCreated, &genero)
}

// @BasePath /
//
// @Summary Atualiza genero
// @Schemes
// @Description Atualiza os dados de um genero
// @Tags generos
// @Accept json
// @Produce json
// @Param id path int true "ID do genero"
// @Param genero body models.Genero true "Dados do genero"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /generos/{id} [put]
func UpdateGenero(c *gin.Context) {
	var genero, generoInput models.Genero
	models.Get(c, database.DB, &genero)

	if err := c.ShouldBindJSON(&generoInput); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	genero.Update(c, database.DB, &generoInput)
	c.JSON(http.StatusNoContent, nil)
}

// @BasePath /
//
// @Summary Deleta genero
// @Schemes
// @Description Remove o cadastro de um genero
// @Tags generos
// @Accept json
// @Produce json
// @Param id path int true "ID do genero"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /generos/{id} [delete]
func DeleteGenero(c *gin.Context) {
	models.Delete[models.Genero](c, database.DB)
	c.JSON(http.StatusNoContent, nil)
}
