package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	// r.GET("/ping", pingHandler)
	// r.POST("/ping", pingPostHandler)
	// r.GET("/students", getStudentHandler)
	// r.POST("/students", postStudentHandler)

	r.GET("/api/todos", getTodosHandler)
	r.GET("/api/todos/:id", getTodosByIdHandler)
	r.POST("/api/todos", postTodosHandler)
	r.DELETE("/api/todos/:id", deleteTodosByIdHandler)

	//Add Commit

	r.Run(":1234")

}
