package tests
 
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"todoapp/database"
	"todoapp/models"
	"todoapp/routes"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setUpRouter() *gin.Engine {
	router := gin.Default()
	routes.SetupRoutes(router)
	return router
}

func TestCreateTask(t *testing.T) {
	database.SetupDatabase()
	router := setUpRouter()
	task := models.Task {
		Title: "unit test1 task",
		Description: "This is a test task..",
	}
	// gives me json format of given struct instance
	taskJSON , _ := json.Marshal(task)
	w := httptest.NewRecorder()
	req , _ := http.NewRequest("POST" , "/tasks" , bytes.NewBuffer(taskJSON))
	req.Header.Set("Content-Type" , "application/json")
	router.ServeHTTP(w , req);
	assert.Equal(t , http.StatusCreated , w.Code)
	var responseTask models.Task
	json.Unmarshal(w.Body.Bytes(), &responseTask)
	assert.Equal(t, "unit test1 task", responseTask.Title)
}

func TestGetTasks(t *testing.T) {
	database.SetupDatabase()
	router := setUpRouter()
	w := httptest.NewRecorder()
	req , _ := http.NewRequest("GET" , "/tasks" , nil)
	router.ServeHTTP(w , req)
	assert.Equal(t , http.StatusOK , w.Code);
}

func TestGetTaskById(t *testing.T) {
	database.SetupDatabase()
	router := setUpRouter()
	w := httptest.NewRecorder()
	req , _ := http.NewRequest("GET" , "/tasks/1" , nil)
	router.ServeHTTP(w , req)
	assert.Equal(t , http.StatusOK, w.Code);
}

func TestDeleteTask(t *testing.T) {
	database.SetupDatabase()
	router := setUpRouter()
	w := httptest.NewRecorder()
	req , _ := http.NewRequest("DELETE" , "tasks/1" , nil)
	router.ServeHTTP(w , req)
	assert.Equal(t , http.StatusNotFound , w.Code)
}