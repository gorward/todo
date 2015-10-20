package app

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TodoListInterface interface for TodoListRepositories
type TodoListInterface interface {
	AddTodo(string, Todo) error
	RemoveTodo(string, string) error
	UpdateTodo(string, Todo) error
	FetchTodoLists() ([]TodoList, error)
	FetchTodoListsByID(string) ([]TodoList, error)
}

// TodoList struct
type TodoList struct {
	ID         string          `bson:"_id,omitempty" json:"-"`
	TodoListID string          `bson:"todo_list_id,omitempty" json:"todo_list_id"`
	Name       string          `bson:"name,omitempty" json:"name"`
	Todos      map[string]Todo `bson:"todos,omitempty" json:"todos"`
}

// Todo struct
type Todo struct {
	TodoID      string `bson:"todo_id,omitempty" json:"todo_id"`
	Name        string `bson:"name,omitempty" json:"name"`
	Description string `bson:"description,omitempty" json:"description"`
	Status      string `bson:"status,omitempty" json:"status"`
}

// TodoListRepository implements TodoListInterface with mongodb using mgo
type TodoListRepository struct {
	MongoDb *mgo.Database
}

// AddTodo adds new Todo
func (tl TodoListRepository) AddTodo(tlID string, todo Todo) error {
	mgoC := tl.MongoDb.C("todos")

	query := bson.M{"todo_list_id": tlID}
	updateData := bson.M{"$set": bson.M{"todos." + todo.TodoID: todo}}

	_, err := mgoC.Upsert(query, updateData)

	return err
}

// RemoveTodo removes Todo
func (tl TodoListRepository) RemoveTodo(tlID string, tID string) error {
	mgoC := tl.MongoDb.C("todos")

	query := bson.M{"todo_list_id": tlID}
	updateData := bson.M{"$unset": bson.M{"todos." + tID: ""}}

	return mgoC.Update(query, updateData)
}

// UpdateTodo updates Todo
func (tl TodoListRepository) UpdateTodo(tlID string, todo Todo) error {
	mgoC := tl.MongoDb.C("todos")

	prefix := "todos." + todo.TodoID
	update := SetExisting(prefix, StructToBson(&todo))

	query := bson.M{"todo_list_id": tlID}
	updateData := bson.M{"$set": update}

	return mgoC.Update(query, updateData)
}

// FetchTodoLists get all todolist
func (tl TodoListRepository) FetchTodoLists() ([]TodoList, error) {
	todoList := []TodoList{}

	mgoC := tl.MongoDb.C("todos")

	err := mgoC.Find(nil).All(&todoList)

	return todoList, err
}

// FetchTodoListsByID get todolist by id
func (tl TodoListRepository) FetchTodoListsByID(tlID string) ([]TodoList, error) {
	todoList := []TodoList{}

	mgoC := tl.MongoDb.C("todos")

	err := mgoC.Find(bson.M{"todo_list_id": tlID}).All(&todoList)

	return todoList, err
}
