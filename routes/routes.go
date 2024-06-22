package routes

import (
	"todoapp/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/tasks" , controllers.CreateTask);
	router.GET("/tasks" , controllers.GetTasks);
	router.GET("/tasks/:id" , controllers.GetTaskByID);
	router.DELETE("/tasks/:id" , controllers.DeleteTask);
}