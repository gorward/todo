package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorward/mux"
	"net/http"
)

// GetTodoList http.Handler for getting all todo lists
type GetTodoList struct {
	TodoListRepository TodoListInterface
}

func (gtl GetTodoList) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	todoList, err := gtl.TodoListRepository.FetchTodoLists()
	if err != nil {
		RespondErr(w, http.StatusNotFound, err.Error())
		return
	}

	data, err := json.Marshal(todoList)

	RespondData(w, string(data))
}

// GetTodoListID http.Handler for getting todo list with id
type GetTodoListID struct {
	TodoListRepository TodoListInterface
}

func (gtl GetTodoListID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tlID := mux.Vars(r)["id"]
	todoList, err := gtl.TodoListRepository.FetchTodoListsByID(tlID)
	if err != nil {
		RespondErr(w, http.StatusNotFound, err.Error())
		return
	}

	data, err := json.Marshal(todoList)

	RespondData(w, string(data))
}

// CreateTodo creates todo item inside a list
type CreateTodo struct {
	TodoListRepository TodoListInterface
}

func (ct CreateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: validation
	tlID := mux.Vars(r)["tlid"]

	name := r.FormValue("name")
	description := r.FormValue("description")
	status := r.FormValue("status")

	todo := &Todo{
		TodoID:      GenerateUUID(),
		Name:        name,
		Description: description,
		Status:      status,
	}
	fmt.Printf("\n%+v\n", *todo)

	err := ct.TodoListRepository.AddTodo(tlID, *todo)

	if err != nil {
		RespondErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondData(w, "200")
}

// UpdateTodo updates todo item
type UpdateTodo struct {
	TodoListRepository TodoListInterface
}

func (ut UpdateTodo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: validation
	tlID := mux.Vars(r)["tlid"]
	tID := mux.Vars(r)["tid"]

	name := r.FormValue("name")
	description := r.FormValue("description")
	status := r.FormValue("status")

	todo := &Todo{
		TodoID:      tID,
		Name:        name,
		Description: description,
		Status:      status,
	}
	fmt.Printf("\n%+v\n", *todo)

	err := ut.TodoListRepository.UpdateTodo(tlID, *todo)

	if err != nil {
		RespondErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondData(w, "200")
}
