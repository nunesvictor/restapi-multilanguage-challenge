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
// @Summary Lista editoras
// @Schemes
// @Description Recupera a lista de editoras
// @Tags editoras
// @Accept json
// @Produce json
// @Success 200 {array} models.Editora
// @Failure 500 {object} httputil.HTTPError
// @Router /editoras/ [get]
func GetEditoras(c *gin.Context) {
	var editoras []models.Editora

	if result := database.DB.Find(&editoras); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	c.JSON(http.StatusOK, &editoras)
}

// @BasePath /
//
// @Summary Recupera editora
// @Schemes
// @Description Recupera uma editora por id
// @Tags editoras
// @Accept json
// @Produce json
// @Param id path int true "ID da editora"
// @Success 200 {object} models.Editora
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /editoras/{id} [get]
func GetEditora(c *gin.Context) {
	var editora models.Editora

	if statusCode, err := findRegister(c, &editora); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	c.JSON(http.StatusOK, &editora)
}

// @BasePath /
//
// @Summary Cria editora
// @Schemes
// @Description Cria uma nova editora
// @Tags editoras
// @Accept json
// @Produce json
// @Param editora body schemas.EditoraInput true "Dados da editora"
// @Success 201 {object} models.Editora
// @Failure 500 {object} httputil.HTTPError
// @Router /editoras/ [post]
func CreateEditora(c *gin.Context) {
	var editora models.Editora

	if err := c.ShouldBindJSON(&editora); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if result := database.DB.Create(&editora); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	if result := database.DB.First(&editora); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	c.JSON(http.StatusCreated, &editora)
}

// @BasePath /
//
// @Summary Atualiza editora
// @Schemes
// @Description Atualiza os dados de uma editora
// @Tags editoras
// @Accept json
// @Produce json
// @Param id path int true "ID da editora"
// @Param editora body schemas.EditoraInput true "Dados da editora"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /editoras/{id} [put]
func UpdateEditora(c *gin.Context) {
	var editora models.Editora
	var editoraInput schemas.EditoraInput

	if statusCode, err := findRegister(c, &editora); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	if err := c.ShouldBindJSON(&editoraInput); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	editora.UpdateFromInput(&editoraInput)

	if result := database.DB.Save(&editora); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @BasePath /
//
// @Summary Deleta editora
// @Schemes
// @Description Remove o cadastro de uma editora
// @Tags editoras
// @Accept json
// @Produce json
// @Param id path int true "ID da editora"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /editoras/{id} [delete]
func DeleteEditora(c *gin.Context) {
	var editora models.Editora

	if statusCode, err := findRegister(c, &editora); err != nil {
		httputil.NewError(c, statusCode, err)
		return
	}

	if result := database.DB.Delete(&editora); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
