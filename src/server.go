package main

import (
	//	"fmt"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {

	log.Println("start")

	os.Remove("./foo.db")

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
