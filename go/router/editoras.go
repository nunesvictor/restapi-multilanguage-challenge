package router

import (
	"gin-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func EditorasGroupRouter(r *gin.Engine) {
	editoras := r.Group("/editoras")
	{
		editoras.GET("/", controllers.GetEditoras)
		editoras.POST("/", controllers.CreateEditora)
		editoras.GET("/:id", controllers.GetEditora)
		editoras.PUT("/:id", controllers.UpdateEditora)
		editoras.DELETE("/:id", controllers.DeleteEditora)
	}
}
