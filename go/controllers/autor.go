package controllers

import (
	"gin-restapi/database"
	"gin-restapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAutores(ctx *gin.Context) {
	var autores []models.Autor
	database.DB.Find(&autores)

	ctx.JSON(http.StatusOK, autores)
}
