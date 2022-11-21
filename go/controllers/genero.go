package controllers

import (
	"fmt"
	"gin-restapi/database"
	"gin-restapi/httputil"
	"gin-restapi/models"
	"gin-restapi/schemas"
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

	if result := database.DB.Find(&generos); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

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

	if statusCode, err := findRegister(c, &genero); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

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
// @Param genero body schemas.GeneroInput true "Dados do genero"
// @Success 201 {object} models.Genero
// @Failure 500 {object} httputil.HTTPError
// @Router /generos/ [post]
func CreateGenero(c *gin.Context) {
	var genero models.Genero

	if err := c.ShouldBindJSON(&genero); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if result := database.DB.Create(&genero); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	if result := database.DB.First(&genero); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

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
// @Param genero body schemas.GeneroInput true "Dados do genero"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /generos/{id} [put]
func UpdateGenero(c *gin.Context) {
	var genero models.Genero
	var generoInput schemas.GeneroInput

	if statusCode, err := findRegister(c, &genero); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	if err := c.ShouldBindJSON(&generoInput); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	genero.UpdateFromInput(&generoInput)

	if result := database.DB.Save(&genero); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

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
	var genero models.Genero

	if statusCode, err := findRegister(c, &genero); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	if result := database.DB.Delete(&genero); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
