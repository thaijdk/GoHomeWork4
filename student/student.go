package student

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Student struct {
	ID   int    `json:"student_id"`
	Name string `json:"name"`
}

var students = map[int]Student{
	1: Student{Name: "Sup01", ID: 1},
}

func (s Student) getStudentHandler(c *gin.Context) {
	student := []Student{}
	for _, s := range students {
		student = append(student, s)
	}
	c.JSON(http.StatusOK, student)
}

func (s Student) postStudentHandler(c *gin.Context) {
	s := Student{}
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	id := len(students)
	id++
	s.ID = id
	students[id] = s
	c.JSON(http.StatusOK, students)
}
