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
	router.Get("/", fs)

	router.Group("/api", func(router mux.Router) {
		router.Get("/todolist", getTodoListHandler)
		router.Get("/todolist/{id}", getTodoListIDHandler)
		router.Put("/todolist/{tlid}/todo", createTodoHandler)
		router.Patch("/todolist/{tlid}/todo/{tid}", updateTodoHandler)
		router.Delete("/todolist/{tlid}/todo/{tid}", deleteTodoHandler)
	})

	router.ServeHTTP(w, r)
}
