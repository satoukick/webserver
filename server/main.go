package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satoukick/webserver/config"
	logs "github.com/satoukick/webserver/log"
	"github.com/satoukick/webserver/model"
)

var db *gorm.DB

func init() {
	config.Init()
	initPostgres()
}

func initPostgres() {
	var err error
	pgconf := config.Conf.GetPGEnvString()
	db, err = gorm.Open("postgres", pgconf)
	if err != nil {
		logs.Fatal(err)
	}
	db.AutoMigrate(&model.TodoModel{})
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	return router
}

// TODO : do something with goroutine, docker
func main() {
	defer logs.Sync()
	router := setupRouter()
	router.RunTLS(":8080", "../https/server.pem", "../https/server.key")
	// router.Run()
}

func createTodo(c *gin.Context) {
	title := c.PostForm("title")
	if title == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "title should not be empty.",
		})
		return
	}

	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := model.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Todo item created successfully!",
		"resourceId": todo.ID,
	})
}

// fetcHallTodo fetches all records
func fetchAllTodo(c *gin.Context) {
	var todos []model.TodoModel
	var _todos []model.TransformedTodo

	db.Find(&todos)
	if len(todos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo found!",
		})
		return
	}
	for _, item := range todos {
		var completed bool
		if item.Completed == 1 {
			completed = true
		}
		new := model.TransformedTodo{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		}
		_todos = append(_todos, new)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todos,
	})
}

// fetchSingleTodo fetches single records using ID
func fetchSingleTodo(c *gin.Context) {
	todo := model.TodoModel{}
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "todo id not found",
		})
		return
	}

	completed := todo.Completed == 1
	trans := model.TransformedTodo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: completed,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   trans,
	})
}

func updateTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)

	if todo.ID <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo found!",
		})
		return
	}

	completed, _ := strconv.Atoi(c.PostForm("completed"))
	title := c.PostForm("title")
	db.Model(&todo).Updates(model.TodoModel{Completed: completed, Title: title})
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated successfully!",
	})
}

func deleteTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")

	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo found!",
		})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully!",
	})
}
