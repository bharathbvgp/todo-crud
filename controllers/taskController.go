package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todoapp/database"
	"todoapp/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	// creating a new task
	var newTask models.Task
	// binding incoming json data with the new task and handling errors if any occurs
	if err := c.ShouldBindJSON(&newTask) ; err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
		return 
	}	
	// creating a new task in database
	if result := database.DB.Create(&newTask); result.Error != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : result.Error.Error()})
		return 
	}
	// the data has been updated to database
	c.JSON(http.StatusCreated , newTask)

}

func GetTasks(c *gin.Context) {
	// get all tasks
	var tasks []models.Task

	if result := database.DB.Find(&tasks); result.Error != nil {
		c.JSON(http.StatusInternalServerError , gin.H{"error" : result.Error.Error()})
		return 
	}
	c.JSON(http.StatusOK , tasks);
}

func GetTaskByID(c *gin.Context) {
	fmt.Println("entered ... ...... .....")
	// getting id from params
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task models.Task
	// fetching the task from the database
	if result := database.DB.First(&task, id); result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
	fmt.Println("task is : " ,task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	// getting the id from database
	id , err := strconv.Atoi(c.Param("id")); 
	if err != nil {
		c.JSON(http.StatusBadRequest , gin.H{"error" : "Invalid task ID"})
		return 
	}
	// deleting task from database
	if result := database.DB.Delete(&models.Task{} , id) ; result.Error != nil {
		if result.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound , gin.H{"error" : "Task not found"})
		} else {
			c.JSON(http.StatusInternalServerError , gin.H{"error" : result.Error.Error()})
		}
		return
	}
	// sending response with status code that data is deleted 
	c.JSON(http.StatusNoContent , nil);
}

