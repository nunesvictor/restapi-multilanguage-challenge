package router

import (
	"github.com/gin-gonic/gin"

	docs "gin-restapi/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(setupDB func()) *gin.Engine {
	setupDB()

	r := gin.New()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	AutoresGroupRouter(r)
	EditorasGroupRouter(r)
	GenerosGroupRouter(r)
	LivrosGroupRouter(r)

	return r
}
