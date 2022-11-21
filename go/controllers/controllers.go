package controllers

import (
	"fmt"
	"gin-restapi/database"
	"gin-restapi/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func findRegister[T any](c *gin.Context, model T) (int, error) {
	var uri schemas.SingleResourceUri

	if err := c.ShouldBindUri(&uri); err != nil {
		return http.StatusBadRequest, fmt.Errorf("invalid path parameter: id")
	}

	if result := database.DB.First(&model, uri.ID); result.Error != nil {
		return http.StatusNotFound, fmt.Errorf("register not found")
	}

	return http.StatusOK, nil
}
