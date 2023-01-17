package main

import (
	handlers "delosTest/Handlers"
	services "delosTest/Services"
	"delosTest/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"time"

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

func library(dueDate, submitDate time.Time) (result int) {
	lateRange := submitDate.YearDay() - dueDate.YearDay()
	if lateRange <= 0 || dueDate.Year() > submitDate.Year() {
		result = 0
		return
	} else if submitDate.Year() != dueDate.Year() {
		result = 12000
		return
	} else if submitDate.Month() != dueDate.Month() {
		result = (int(submitDate.Month()) - int(dueDate.Month())) * 500
		return
	} else {
		result = 15 * lateRange
		return
	}
}

func whoGetSourCandy(student, candies, firstStudent int) (result int) {
	result = (candies + firstStudent - 1) % student
	if result == 0 {
		result = student
	}
	return
}

func sameSumElement(arrLen int, arr []int) (result string) {
	for index := range arr {
		left := 0
		right := 0

		for i := 0; i < arrLen; i++ {
			if i > index {
				right += arr[i]
			}
			if i < index {
				left += arr[i]
			}
		}
		if right == left {
			result = "YES"
			return
		}
	}
	result = "NO"
	return
}
