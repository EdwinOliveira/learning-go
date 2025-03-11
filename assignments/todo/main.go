package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	id uint32
	name string
	description string
	isCompleted bool
	isActive bool
}

var todos = []Todo{}

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
				fmt.Println("Todo ID?")
				todoId, _ := reader.ReadString('\n')
				trimmedTodoId := strings.TrimSpace(todoId)
				parsedTodoId, parsedTodoIdError := strconv.Atoi(strings.TrimSpace(trimmedTodoId))

				if parsedTodoIdError != nil {
					fmt.Println("Error converting string to integer:", parsedTodoIdError)
					return
				}

				foundTodo, foundTodoError := Todo.findTodoById(Todo{id: uint32(parsedTodoId)})

				if foundTodoError != nil {
					fmt.Println("Error finding the todo:", foundTodoError)
					return
				}

				fmt.Println(foundTodo)
				break
			}
			case "2": {
				fmt.Println("Todo Name?")
				todoName, _ := reader.ReadString('\n')
				trimmedTodoName := strings.TrimSpace(todoName)

				fmt.Println("Todo Description?")
				todoDescription, _ := reader.ReadString('\n')
				trimmedTodoDescription := strings.TrimSpace(todoDescription)

				createdTodoResponse := Todo.createTodo(Todo{ name: trimmedTodoName, description: trimmedTodoDescription })

				if createdTodoResponse == 1 {
					fmt.Println("Todo created with success")
				}

				break
			}
			case "3": {
				fmt.Println("Todo ID?")
				todoId, _ := reader.ReadString('\n')
				trimmedTodoId := strings.TrimSpace(todoId)
				parsedTodoId, parsedTodoIdError := strconv.Atoi(strings.TrimSpace(trimmedTodoId))

				if parsedTodoIdError != nil {
					fmt.Println("Error converting string to integer:", parsedTodoIdError)
					return
				}

				fmt.Println("Todo Name?")
				todoName, _ := reader.ReadString('\n')
				trimmedTodoName := strings.TrimSpace(todoName)

				fmt.Println("Todo Description?")
				todoDescription, _ := reader.ReadString('\n')
				trimmedTodoDescription := strings.TrimSpace(todoDescription)

				_, updatedTodoResponseError := Todo.updateTodoById(Todo{ id: uint32(parsedTodoId), name: trimmedTodoName, description: trimmedTodoDescription })

				if updatedTodoResponseError != nil {
					fmt.Println(updatedTodoResponseError)
					return
				}

				fmt.Println("Todo updated with sucess!")
				break
			}
			case "4": {
				fmt.Println("Todo ID?")
				todoId, _ := reader.ReadString('\n')
				trimmedTodoId := strings.TrimSpace(todoId)
				parsedTodoId, parsedTodoIdError := strconv.Atoi(strings.TrimSpace(trimmedTodoId))

				if parsedTodoIdError != nil {
					fmt.Println("Error converting string to integer:", parsedTodoIdError)
					return
				}

				_, updatedTodoResponseError := Todo.updateTodoById(Todo{ id: uint32(parsedTodoId), isCompleted: true })

				if updatedTodoResponseError != nil {
					fmt.Println(updatedTodoResponseError)
					return
				}

				fmt.Println("Todo updated with sucess!")
				break
			}
			case "5": {
				fmt.Println("Todo ID?")
				todoId, _ := reader.ReadString('\n')
				trimmedTodoId := strings.TrimSpace(todoId)
				parsedTodoId, parsedTodoIdError := strconv.Atoi(strings.TrimSpace(trimmedTodoId))

				if parsedTodoIdError != nil {
					fmt.Println("Error converting string to integer:", parsedTodoIdError)
					return
				}

				_, updatedTodoResponseError := Todo.updateTodoById(Todo{ id: uint32(parsedTodoId), isActive: false })

				if updatedTodoResponseError != nil {
					fmt.Println(updatedTodoResponseError)
					return
				}

				fmt.Println("Todo updated with sucess!")
				break
			}
		}
	}
}
