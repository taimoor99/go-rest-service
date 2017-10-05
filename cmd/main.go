package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedrocelso/go-rest-service/lib/controllers"
)

func init() {
	router := gin.New()
	router.Use(controllers.CORSMiddleware())
	router.Use(controllers.SetContextValues())
	router.OPTIONS("/", func(c *gin.Context) {})

	v1 := router.Group("/v1")

	users := v1.Group("/users")
	users.POST("/", controllers.CreateUser)
	users.GET("/:userEmail", controllers.GetUser)
	users.GET("/", controllers.GetUsers)
	users.PUT("/:userEmail", controllers.UpdateUser)
	users.DELETE("/:userEmail", controllers.DeleteUser)

	tasks := v1.Group("/tasks")
	tasks.POST("/", controllers.CreateTask)
	tasks.GET("/:taskId", controllers.GetTask)

	http.Handle("/", router)
}
