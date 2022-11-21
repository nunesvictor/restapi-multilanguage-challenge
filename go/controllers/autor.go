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
// @Summary Lista autores
// @Schemes
// @Description Recupera a lista de autores
// @Tags autores
// @Accept json
// @Produce json
// @Success 200 {array} models.Autor
// @Router /autores/ [get]
func GetAutores(c *gin.Context) {
	var autores []models.Autor
	database.DB.Find(&autores)
	c.JSON(http.StatusOK, &autores)
}

// @BasePath /
//
// @Summary Recupera autor
// @Schemes
// @Description Recupera um autor por id
// @Tags autores
// @Accept json
// @Produce json
// @Param id path int true "ID do autor"
// @Success 200 {object} models.Autor
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /autores/{id} [get]
func GetAutor(c *gin.Context) {
	var autor models.Autor

	if statusCode, err := findRegister(c, &autor); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	c.JSON(http.StatusOK, &autor)
}

// @BasePath /
//
// @Summary Cria autor
// @Schemes
// @Description Cria um novo autor
// @Tags autores
// @Accept json
// @Produce json
// @Param autor body schemas.AutorInput true "Dados do autor"
// @Success 201 {object} models.Autor
// @Failure 500 {object} httputil.HTTPError
// @Router /autores/ [post]
func CreateAutor(c *gin.Context) {
	var autor models.Autor
	c.ShouldBindJSON(&autor)

	if result := database.DB.Create(&autor); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	database.DB.First(&autor, autor.ID)
	c.JSON(http.StatusCreated, &autor)
}

// @BasePath /
//
// @Summary Atualiza autor
// @Schemes
// @Description Atualiza os dados de um autor
// @Tags autores
// @Accept json
// @Produce json
// @Param id path int true "ID do autor"
// @Param autor body schemas.AutorInput true "Dados do autor"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /autores/{id} [put]
func UpdateAutor(c *gin.Context) {
	var autor models.Autor
	var autorInput schemas.AutorInput

	c.ShouldBindJSON(&autorInput)

	if statusCode, err := findRegister(c, &autor); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	autor.UpdateFromInput(&autorInput)
	database.DB.Save(&autor)

	c.JSON(http.StatusNoContent, nil)
}

// @BasePath /
//
// @Summary Deleta autor
// @Schemes
// @Description Remove o cadastro de um autor
// @Tags autores
// @Accept json
// @Produce json
// @Param id path int true "ID do autor"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /autores/{id} [delete]
func DeleteAutor(c *gin.Context) {
	var autor models.Autor

	if statusCode, err := findRegister(c, &autor); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	database.DB.Delete(&autor)
	c.JSON(http.StatusNoContent, nil)
}
