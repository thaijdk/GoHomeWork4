package student

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/thaijdk/GoHomeWork4/database"
)

type Student struct {
	ID   int    `json:"student_id"`
	Name string `json:"name"`
}

var students = map[int]Student{
	1: Student{Name: "Sup01", ID: 1},
}

func (s Student) GetHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	stmt, _ := db.Prepare("SELECT id, name FROM todos")

	students := []Student{}

	rows, _ := stmt.Query()
	for rows.Next() {

		s := Student{}

		err := rows.Scan(&s.ID, &s.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		students = append(students, s)
	}
	fmt.Println(students)

	c.JSON(http.StatusOK, students)
}

func (s Student) GetByIdHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	stmt, _ := db.Prepare("SELECT id, name FROM todos WHERE id=$1")

	id := c.Param("id")

	row := stmt.QueryRow(id)

	s = Student{}

	err2 := row.Scan(&s.ID, &s.Name)
	if err2 != nil {
		log.Fatal("error", err.Error())
	}

	c.JSON(http.StatusOK, s)
}

func (s Student) PostHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	studentVal := Student{}
	if err := c.ShouldBindJSON(&studentVal); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if studentVal.Name != "" {
		query := `
	INSERT INTO student (name) VALUES ($1) RETURNING id
	`
		var id int
		row := db.QueryRow(query, studentVal.Name)
		err = row.Scan(&id)
		if err != nil {
			log.Fatal("can't scan id", err.Error())
		}

		c.JSON(http.StatusOK, "Insert Completed")
	} else {
		c.JSON(http.StatusOK, "Empty Input")
	}
}

func (s Student) DeleteByIdHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM student WHERE id=$1")

	id := c.Param("id")

	if _, err := stmt.Exec(id); err != nil {
		log.Fatal("exec error ", err.Error())
	}

	c.JSON(http.StatusOK, "delete complete")
}
