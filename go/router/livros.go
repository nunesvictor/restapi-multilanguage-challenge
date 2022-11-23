package router

import (
	"gin-restapi/controllers"

	"github.com/gin-gonic/gin"
)

func LivrosGroupRouter(r *gin.Engine) {
	livros := r.Group("/livros")
	{
		livros.GET("/", controllers.GetLivros)
		livros.POST("/", controllers.CreateLivro)
		livros.GET("/:id", controllers.GetLivro)
		livros.PUT("/:id", controllers.UpdateLivro)
		livros.DELETE("/:id", controllers.DeleteLivro)
	}
}
