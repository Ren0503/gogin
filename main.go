package main

import (
	"fmt"
	"os"
	"ren0503/gogin/controllers"
	"ren0503/gogin/infrastructure"
	"ren0503/gogin/models"
	"ren0503/gogin/seeds"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Todo{})
}

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}

	database := infrastructure.OpenDbConnection()
	defer database.Close()
	migrate(database)
	seeds.Seed(database)

	goGonicEngine := gin.Default()
	goGonicEngine.Use(cors.Default())
	goGonicEngine.GET("/api/todos", controllers.GetAllTodos).
		// GET("/api/todos/completed", controllers.GetAllPendingTodos).
		// GET("/api/todos/pending", controllers.GetAllCompletedTodos).
		GET("/api/todos/:id", controllers.GetTodoById)

	// This is how you should do it, the above was just to get started :)
	apiGroup := goGonicEngine.Group("/api")
	apiGroup.POST("/todos", controllers.CreateTodo)
	apiGroup.PUT("/todos/:id", controllers.UpdateTodo)
	apiGroup.PATCH("/todos/:id", controllers.CreateTodo)

	apiGroup.DELETE("/todos", controllers.DeleteAllTodos)
	apiGroup.DELETE("/todos/:id", controllers.DeleteTodo)

	goGonicEngine.Run(":8080")
}
