package routes

import (
	"github.com/2gbeh/drygin-api/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(controller *controllers.TagController) *gin.Engine {
	route := gin.Default()
	// swagger
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// tags
	apiRoute := route.Group("/api")
	tagRoutes := apiRoute.Group("/tags")
	tagRoutes.GET("", controller.All)
	tagRoutes.GET("/migrate", controller.Migrate)
	tagRoutes.GET("/migrate-fresh", controller.MigrateFresh)
	tagRoutes.GET("/migrate-fresh-seed", controller.MigrateFreshSeed)

	return route
}
