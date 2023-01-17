package main

import (
	handlers "delosTest/Handlers"
	services "delosTest/Services"
	"delosTest/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	libraryService := services.NewLibraryService()
	candyService := services.NewCandyService()
	arrayElementService := services.NewArrayElementService()

	libraryHandler := handlers.NewLibraryHandler(libraryService)
	candyHandler := handlers.NewCandyHandler(candyService)
	arrayElementHandler := handlers.NewArrayElementHandler(arrayElementService)
	router := gin.Default()
	api := router.Group("api/v1")
	docs.SwaggerInfo.BasePath = "/api/v1"

	api.GET("/library/bookreturn/:dueDate/:submitDate", libraryHandler.BookReturn)
	api.GET("/candy/whogetsourcandy/:student/:candies/:firstStudent", candyHandler.WhoGetSourCandy)
	api.GET("/arrayelement/samesumelement/:arrLen/:arr", arrayElementHandler.SameSumElement)
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
