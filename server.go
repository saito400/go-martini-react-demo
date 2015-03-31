package main

import (
	"database/sql"
	_ "fmt"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	_ "github.com/mattn/go-sqlite3"
	"log"
	_ "net/http"
	_ "os"
	_ "strconv"
)

func main() {

	log.Println("main start")

	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/todo/list.json", func(r render.Render) {

		r.JSON(200, list())

	})

	m.Get("/hello", func(r render.Render) {
		r.HTML(200, "hello", "you")
	})

	m.Post("/todo/create", binding.Bind(Todo{}), func(r render.Render, todo Todo) {
		insert(todo)
		r.JSON(200, list())

	})

	m.Run()

}

func insert(todo Todo) {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into todo(id, content) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(todo.Id, todo.Content)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

}

func list() []Todo {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select id, content from todo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todoList []Todo

	for rows.Next() {
		var id int
		var content string
		rows.Scan(&id, &content)
		todoList = append(todoList, Todo{
			id, content,
		})

	}

	return todoList
}

type Todo struct {
	Id      int    `form:"id" json:"id" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}
