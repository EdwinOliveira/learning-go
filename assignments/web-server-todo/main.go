package main

import (
	"fmt"
	"net/http"
)

var todos = []Todo{}

type Todo struct {
	id uint32
	name string
	description string
	isCompleted bool
	isActive bool
}

func (t Todo) findTodoById() (Todo, error) {
	for _, todo := range todos {
		if todo.id == t.id {
			return todo, nil
		}
	}

	return Todo{}, fmt.Errorf("No Todo found with that ID!")
}

func (t Todo) createTodo() int16 {
	todos = append(todos, Todo{ 
		id: uint32(len(todos) + 1), 
		name: t.name, 
		description: t.description, 
		isCompleted: false, 
		isActive: true,
	})

	return 1
}

func (t Todo) updateTodoById() (int16, error) {
	todoIndex := -1

	for i, todo := range todos {
		if todo.id == t.id {
			todoIndex = i 
		}
	}

	if todoIndex == -1 {
		return -1, fmt.Errorf("No Todo found with that ID!")
	}
 
	if t.name != "" && t.name != todos[todoIndex].name {
		todos[todoIndex].name = t.name
	}

	if t.description != "" && t.description != todos[todoIndex].description {
		todos[todoIndex].description = t.description
	}

	if t.isCompleted != todos[todoIndex].isCompleted {
		todos[todoIndex].isCompleted = t.isCompleted
	}

	if t.isActive != todos[todoIndex].isActive {
		todos[todoIndex].isActive = t.isActive
	}

	return 1, nil
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet: {}
			case http.MethodPost: {}
			case http.MethodPut: {}
		}
	})

	http.ListenAndServe(":8090", nil)
}