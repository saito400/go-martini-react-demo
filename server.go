package main

import (
	"database/sql"
	_ "fmt"
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

	//	m.Post("/todo", func(r render.Render) {
	//		add()
	//		r.JSON(200, list())
	//	})

	m.Run()

}

//func add(r *http.Request) {
//
//	id, _ := strconv.Atoi(r.FormValue("id"))
//	content := r.FormValue("content")
//
//	var todo = Todo{
//		Id:      id,
//		Content: content,
//	}
//}

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
	Id      int    `json:"id"`
	Content string `json:"content"`
}
