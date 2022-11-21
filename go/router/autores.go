package router

import (
	"gin-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func AutoresGroupRouter(r *gin.Engine) {
	autores := r.Group("/autores")
	{
		autores.GET("/", controllers.GetAutores)
		autores.POST("/", controllers.CreateAutor)
		autores.GET("/:id", controllers.GetAutor)
		autores.PUT("/:id", controllers.UpdateAutor)
		autores.DELETE("/:id", controllers.DeleteAutor)
	}
}
