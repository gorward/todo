package main

import (
	"github.com/gorward/todo/app"
	"gopkg.in/mgo.v2"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	mdbSession, _ := mgo.Dial("127.0.0.1")
	tlr := app.TodoListRepository{
		MongoDb: mdbSession.DB("todos"),
	}

	http.ListenAndServe(":8080", app.Handler{
		TodoListRepository: tlr,
	})

	defer mdbSession.Close()
}
