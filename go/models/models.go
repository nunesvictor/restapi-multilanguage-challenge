package models

import (
	"fmt"
	"gin-restapi/httputil"
	"gin-restapi/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create[T any](c *gin.Context, db *gorm.DB, model T) {
	if err := c.ShouldBindJSON(&model); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	if result := db.Create(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	if result := db.First(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}
}

func Delete[T any](c *gin.Context, db *gorm.DB) {
	var uri schemas.SingleResourceUri
	var model T

	if err := c.ShouldBindUri(&uri); err != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("invalid path parameter: id"))
		return
	}

	if result := db.Delete(&model, uri.ID); result.Error != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("not found"))
		return
	}
}

func Get[T any](c *gin.Context, db *gorm.DB, model T) {
	var uri schemas.SingleResourceUri

	if err := c.ShouldBindUri(&uri); err != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("invalid path parameter: id"))
		return
	}

	if result := db.First(&model, uri.ID); result.Error != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("not found"))
		return
	}
}

func ListAll[T any](c *gin.Context, db *gorm.DB, model T) {
	if result := db.Find(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}
}

func Update[T any](c *gin.Context, db *gorm.DB, model T, postSave func()) {
	if result := db.Save(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return
	}

	if postSave != nil {
		postSave()
	}
}
