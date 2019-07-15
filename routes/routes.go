package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	Controller "miniProject/controllers"
)

func InitiateRoutes()  {
	router := gin.Default()

	v1 := router.Group("/terminal")
	{
		v1.POST("/", Controller.AddTerminal)
		v1.GET("/", Controller.FetchTerminals)
		//v1.GET("/:id", fetchSingleTodo)
		//v1.PUT("/:id", updateTodo)
		//v1.DELETE("/:id", deleteTodo)
	}

	fmt.Println("firing routes")

	router.Run()
}

