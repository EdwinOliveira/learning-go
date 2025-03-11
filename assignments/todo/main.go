package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type Todo struct {
	id uint32
	name string
	description string
	isCompleted bool
	isActive bool
}

type TodoRepository interface {
	findTodoById(id uint32) (Todo, error)
	createTodo(name, description string) int16
	updateTodoById(id uint32, name, description string, isCompleted, isActive bool) (int16, error)
}

var todos = []Todo{}

func (t Todo) findTodoById(id uint32) (Todo, error) {
	for _, todo := range todos {
		if todo.id == id {
			return todo, nil
		}
	}

	return Todo{}, fmt.Errorf("No Todo found with that ID!")
}

func (t Todo) createTodo(name, description string) int16 {
	todos = append(todos, Todo{ 
		id: uint32(len(todos) + 1), 
		name: name, 
		description: description, 
		isCompleted: false, 
		isActive: true,
	})

	return 1
}

func (t Todo) updateTodoById(id uint32, name, description string, isCompleted, isActive bool) (int16, error) {
	todoIndex := -1

	for i, todo := range todos {
		if todo.id == id {
			todoIndex = i 
		}
	}

	if todoIndex == -1 {
		return -1, fmt.Errorf("No Todo found with that ID!")
	}
 
	if name != "" && name != todos[todoIndex].name {
		todos[todoIndex].name = name
	}

	if description != "" && description != todos[todoIndex].description {
		todos[todoIndex].description = description
	}

	if isCompleted != todos[todoIndex].isCompleted {
		todos[todoIndex].isCompleted = isCompleted
	}

	if isActive != todos[todoIndex].isActive {
		todos[todoIndex].isActive = isActive
	}

	return 1, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Hello. What's the action?")
		fmt.Println("1 - Find Todo")
		fmt.Println("2 - Create Todo")
		fmt.Println("3 - Update Todo")
		fmt.Println("4 - Complete Todo")
		fmt.Println("5 - Remove Todo")
		action, _ := reader.ReadString('\n')
		trimmedAction := strings.TrimSpace(action)

		switch(trimmedAction) {
			case "1": {
				break
			}
			case "2": {
				break
			}
			case "3": {
				break
			}
			case "4": {
				break
			}
			case "5": {
				break
			}
		}
	}
}
