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

	models.ListAll(c, database.DB, &editoras)
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

	models.Get(c, database.DB, &editora)
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

	models.Create(c, database.DB, &editora)
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
	var editora, editoraInput models.Editora
	models.Get(c, database.DB, &editora)

	if err := c.ShouldBindJSON(&editoraInput); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	editora.Update(c, database.DB, &editoraInput)
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
	models.Delete[models.Editora](c, database.DB)
	c.JSON(http.StatusNoContent, nil)
}
