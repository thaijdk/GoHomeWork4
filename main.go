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
	r.GET("/api/student", s.GetHandler)
	r.GET("/api/student/:id", s.GetByIdHandler)
	r.POST("/api/student", s.PostHandler)
	r.DELETE("/api/student/:id", s.DeleteByIdHandler)

	t := todo.Todo{}
	r.GET("/api/todos", t.GetHandler)
	r.GET("/api/todos/:id", t.GetByIdHandler)
	r.POST("/api/todos", t.PostHandler)
	r.DELETE("/api/todos/:id", t.DeleteByIdHandler)

	r.Run(":1234")

}
