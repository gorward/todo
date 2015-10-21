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
	deleteTodoHandler := DeleteTodo{
		TodoListRepository: h.TodoListRepository,
	}

	router := mux.NewRouter(nil)

	fs := http.FileServer(http.Dir("public"))
	router.GET("/", fs)

	router.GROUP("/api", func(router mux.Router) {
		router.GET("/todolist", getTodoListHandler)
		router.GET("/todolist/{id}", getTodoListIDHandler)
		router.PUT("/todolist/{tlid}/todo", createTodoHandler)
		router.PATCH("/todolist/{tlid}/todo/{tid}", updateTodoHandler)
		router.DELETE("/todolist/{tlid}/todo/{tid}", deleteTodoHandler)
	})

	router.ServeHTTP(w, r)
}
