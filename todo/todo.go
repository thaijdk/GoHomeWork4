package todo

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thaijdk/GoHomeWork4/database"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title`
	Status string `json:"status`
}

func (t Todo) GetHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	stmt, _ := db.Prepare("SELECT id, title, status FROM todos")

	todos := []Todo{}

	rows, _ := stmt.Query()
	for rows.Next() {

		t := Todo{}

		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			log.Fatal(err.Error())
		}
		todos = append(todos, t)
	}
	fmt.Println(todos)

	c.JSON(http.StatusOK, todos)
}

func (t Todo) GetByIdHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	stmt, _ := db.Prepare("SELECT id, title, status FROM todos WHERE id=$1")

	id := c.Param("id")

	row := stmt.QueryRow(id)

	t = Todo{}

	err2 := row.Scan(&t.ID, &t.Title, &t.Status)
	if err2 != nil {
		log.Fatal("error", err.Error())
	}

	c.JSON(http.StatusOK, t)
}

func (t Todo) PostHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	todoVal := Todo{}
	if err := c.ShouldBindJSON(&todoVal); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if (todoVal.Title != "") && (todoVal.Status != "") {
		query := `
	INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id
	`
		var id int
		row := db.QueryRow(query, todoVal.Title, todoVal.Status)
		err = row.Scan(&id)
		if err != nil {
			log.Fatal("can't scan id", err.Error())
		}

		c.JSON(http.StatusOK, "Insert Completed")
	} else {
		c.JSON(http.StatusOK, "Empty Input")
	}
}

func (t Todo) DeleteByIdHandler(c *gin.Context) {
	db, err := sql.Open("postgres", database.Host)
	if err != nil {
		log.Fatal("faltal", err.Error())

	}
	defer db.Close()

	stmt, _ := db.Prepare("DELETE FROM todos WHERE id=$1")

	id := c.Param("id")

	if _, err := stmt.Exec(id); err != nil {
		log.Fatal("exec error ", err.Error())
	}

	c.JSON(http.StatusOK, "delete complete")
}
