package models

import (
	"fmt"
	"gin-restapi/httputil"
	"gin-restapi/schemas"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create[T any](c *gin.Context, db *gorm.DB, model T) error {
	if err := c.ShouldBindJSON(&model); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		log.Printf("aqui")
		return err
	}

	if result := db.Create(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		log.Panicf("%v", result.Error)
		return fmt.Errorf(result.Error.Error())
	}

	if result := db.First(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return fmt.Errorf(result.Error.Error())
	}

	return nil
}

func Delete[T any](c *gin.Context, db *gorm.DB) error {
	var uri schemas.SingleResourceUri
	var model T

	if err := c.ShouldBindUri(&uri); err != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("invalid path parameter: id"))
		return fmt.Errorf("invalid path parameter: id")
	}

	if result := db.Delete(&model, uri.ID); result.Error != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("not found"))
		return fmt.Errorf("not found")
	}

	return nil
}

func Get[T any](c *gin.Context, db *gorm.DB, model T) error {
	var uri schemas.SingleResourceUri

	if err := c.ShouldBindUri(&uri); err != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("invalid path parameter: id"))
		return fmt.Errorf("invalid path parameter: id")
	}

	if result := db.First(&model, uri.ID); result.Error != nil {
		httputil.NewError(c, http.StatusBadRequest, fmt.Errorf("not found"))
		return fmt.Errorf("not found")
	}

	return nil
}

func ListAll[T any](c *gin.Context, db *gorm.DB, model T) error {
	if result := db.Find(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return fmt.Errorf("erro ao processar a requisição")
	}

	return nil
}

func Update[T any](c *gin.Context, db *gorm.DB, model T, postSave func()) error {
	if result := db.Save(&model); result.Error != nil {
		httputil.NewError(c, http.StatusInternalServerError, fmt.Errorf("erro ao processar a requisição"))
		return fmt.Errorf("erro ao processar a requisição")
	}

	if postSave != nil {
		postSave()
	}

	return nil
}
