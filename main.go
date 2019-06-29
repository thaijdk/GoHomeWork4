package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/thaijdk/GoHomeWork4/student"
	"github.com/thaijdk/GoHomeWork4/todo"
)

func authMiddleware(c *gin.Context) {
	fmt.Println("Hello Middlware")
	token := c.GetHeader("Authorization")
	fmt.Println("token:", token)
	if token != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusUnauthorized})
		c.Abort()
		return
	}
	c.Next()
	fmt.Println("Goodbye Middleware")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleware)

	s := student.Student{}
	r.GET("/api/student", s.GetHandler)
	r.GET("/api/student/:id", s.GetByIdHandler)
	r.PUT("/api/student/:id", s.UpdateHandler)
	r.POST("/api/student", s.PostHandler)
	r.DELETE("/api/student/:id", s.DeleteByIdHandler)

	t := todo.Todo{}
	r.GET("/api/todos", t.GetHandler)
	r.GET("/api/todos/:id", t.GetByIdHandler)
	r.PUT("/api/todos/:id", t.UpdateHandler)
	r.POST("/api/todos", t.PostHandler)
	r.DELETE("/api/todos/:id", t.DeleteByIdHandler)

	return r
}
func main() {
	r := setupRouter()
	r.Run(":1234")
}
