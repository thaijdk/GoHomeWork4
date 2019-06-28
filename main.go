package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/thaijdk/GoHomeWork4/student"
	"github.com/thaijdk/GoHomeWork4/todo"
)

func main() {
	r := gin.Default()

	s := student.Student{}
	r.GET("/api/todos", s.GetHandler)
	r.GET("/api/todos/:id", s.GetByIdHandler)
	r.POST("/api/todos", s.PostHandler)
	r.DELETE("/api/todos/:id", s.DeleteByIdHandler)

	t := todo.Todo{}
	r.GET("/api/todos", t.GetHandler)
	r.GET("/api/todos/:id", t.GetByIdHandler)
	r.POST("/api/todos", t.PostHandler)
	r.DELETE("/api/todos/:id", t.DeleteByIdHandler)

	r.Run(":1234")

}
