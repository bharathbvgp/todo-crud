package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"todoapp/database"
	"todoapp/models"
	"todoapp/enums"
	"github.com/gin-gonic/gin"
	"time"
)

// func CreateTask(c *gin.Context) {
// 	// creating a new task
// 	var newTask models.Task
// 	// binding incoming json data with the new task and handling errors if any occurs
// 	if err := c.ShouldBindJSON(&newTask) ; err != nil {
// 		c.JSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
// 		return 
// 	}	
// 	// creating a new task in database
// 	if result := database.DB.Create(&newTask); result.Error != nil {
// 		c.JSON(http.StatusInternalServerError , gin.H{"error" : result.Error.Error()})
// 		return 
// 	}
// 	// the data has been updated to database
// 	c.JSON(http.StatusCreated , newTask)

// }


func CreateTask(c *gin.Context) {

    var newTask models.Task

    if err := c.ShouldBindJSON(&newTask); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newTask.CreatedAt = time.Now() 

    if !isValidPriority(newTask.Priority) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid priority"})
        return
    }
    if !isValidStatus(newTask.Status) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
        return
    }
    if !isValidCategory(newTask.Category) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
        return
    }

    if result := database.DB.Create(&newTask); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusCreated, newTask)
}

func isValidPriority(priority enums.PriorityType) bool {
    switch priority {
    case enums.LowPriority, enums.MediumPriority, enums.HighPriority:
        return true
    default:
        return false
    }
}

func isValidStatus(status enums.StatusType) bool {
    switch status {
    case enums.NewStatus, enums.InProgressStatus, enums.CompletedStatus, enums.CanceledStatus:
        return true
    default:
        return false
    }
}

func isValidCategory(category enums.CategoryType) bool {
    switch category {
    case enums.WorkCategory, enums.PersonalCategory, enums.StudyCategory:
        return true
    default:
        return false
    }
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

