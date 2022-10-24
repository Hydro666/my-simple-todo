package main

import (
	"fmt"
	"log"
	"mytodo/golangtodo/internal"
	"mytodo/golangtodo/server"
)

func main2() {
	app, err := internal.NewApp()
	if err != nil {
		panic(err)
	}

	if err = app.CreateNewList("Foo"); err != nil {
		panic(err)
	}
	if err = app.CreateNewList("Bar"); err != nil {
		panic(err)
	}

	if err = app.AddTaskToList("Foo", "Buy gas"); err != nil {
		panic(err)
	}
	if err = app.AddTaskToList("Foo", "Buy grass"); err != nil {
		panic(err)
	}
	if err = app.AddTaskToList("Foo", "Buy ass"); err != nil {
		panic(err)
	}
	lists, err := app.GetAllListNames()
	if err != nil {
		panic(err)
	}
	fmt.Println(lists)
}

func main() {
	mainServe()
}

func mainServe() {
	ts := server.NewTodoServer()
	log.Println("Starting server")
	log.Fatal(ts.Run())
}
