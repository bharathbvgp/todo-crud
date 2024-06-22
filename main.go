package main

import (
	"todoapp/database"
	"todoapp/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	database.SetupDatabase()
	routes.SetupRoutes(r)
	r.Run()
}