package app

import (
	"github.com/gorward/mux"
	"net/http"
)

// Handler http.Handler of the whole app
type Handler struct {
	TodoListRepository TodoListInterface
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	getTodoListHandler := GetTodoList{
		TodoListRepository: h.TodoListRepository,
	}
	getTodoListIDHandler := GetTodoListID{
		TodoListRepository: h.TodoListRepository,
	}
	createTodoHandler := CreateTodo{
		TodoListRepository: h.TodoListRepository,
	}
	updateTodoHandler := UpdateTodo{
		TodoListRepository: h.TodoListRepository,
	}

	router := mux.NewRouter(nil)

	router.GET("/todolist", getTodoListHandler)
	router.GET("/todolist/{id}", getTodoListIDHandler)
	router.PUT("/todolist/{tlid}/todo", createTodoHandler)
	router.PATCH("/todolist/{tlid}/todo/{tid}", updateTodoHandler)

	router.ServeHTTP(w, r)
}
