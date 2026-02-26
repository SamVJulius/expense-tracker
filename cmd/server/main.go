package main

import (
	"expense-tracker/internal/database"
	"expense-tracker/internal/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	database.LoadEnvVariables()
	database.ConnectDB()
	database.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the Expense Tracker App",
		})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}	

	router.Run()
}
