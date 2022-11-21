package router

import (
	"gin-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func GenerosGroupRouter(r *gin.Engine) {
	generos := r.Group("/generos")
	{
		generos.GET("/", controllers.GetGeneros)
		generos.POST("/", controllers.CreateGenero)
		generos.GET("/:id", controllers.GetGenero)
		generos.PUT("/:id", controllers.UpdateGenero)
		generos.DELETE("/:id", controllers.DeleteGenero)
	}
}
