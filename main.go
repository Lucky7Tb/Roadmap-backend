package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

type todo struct {
	Id          uint16 `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	file, err := os.OpenFile("./todo.json", os.O_CREATE|os.O_WRONLY, 0700)
	if err != nil {
		fmt.Println("Error read json file")
		os.Exit(1)
	}
	defer file.Close()
	existTodos, err := os.ReadFile("./todo.json")
	if err != nil {
		fmt.Println("Error read json file")
		os.Exit(1)
	}

	var todos []todo
	if err := json.Unmarshal(existTodos, &todos); err != nil && err.Error() != "unexpected end of JSON input" {
		fmt.Println("Error parse json to struct")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("Available command: ")
		fmt.Println("1. add [todo]")
		fmt.Println("2. update [todo id] [new todo]")
		fmt.Println("3. delete [todo id]")
		fmt.Println("4. mark-in-progress [todo id]")
		fmt.Println("5. mark-done [todo id]")
		fmt.Println("6. list [done|todo|in-progress]")
		break
	case "add":
		var id uint16 = 1
		if len(todos) > 0 {
			id = todos[len(todos)-1].Id + 1
		}
		newTodo := todo{
			Id:          id,
			Description: os.Args[2],
			Status:      "todo",
			CreatedAt:   time.Now().Format(time.DateTime),
			UpdatedAt:   time.Time{}.Format(time.DateTime),
		}
		todos = append(todos, newTodo)
		json, err := json.Marshal(todos)
		if err != nil {
			fmt.Println("Failed parse struct to json")
		}
		file.Truncate(0)
		if _, err := file.Write(json); err != nil {
			fmt.Println("Failed save todo")
		}
		break
	case "update":
		if len(todos) == 0 {
			fmt.Println("There is no todo to update")
		} else {
			todoId := ""
			todoDescription := ""

			for index := 2; index <= len(os.Args)-1; index++ {
				if index == 2 {
					todoId = os.Args[index]
				}

				if index == 3 {
					todoDescription = os.Args[index]
				}
			}

			if todoId == "" {
				fmt.Println("Please input todo id")
				os.Exit(1)
			}

			if todoDescription == "" {
				fmt.Println("Please input new todo")
				os.Exit(1)
			}

			isExist := false
			for index, todo := range todos {
				if strconv.Itoa(int(todo.Id)) == todoId {
					todos[index].Description = todoDescription
					todos[index].UpdatedAt = time.Now().Format(time.DateTime)
					isExist = true
					break
				}
			}

			if !isExist {
				fmt.Println("Todo not found")
				os.Exit(1)
			}

			json, err := json.Marshal(todos)
			if err != nil {
				fmt.Println("Failed parse struct to json")
			}
			file.Truncate(0)
			if _, err := file.Write(json); err != nil {
				fmt.Println("Failed save todo")
			}
		}
		break
	case "delete":
		if len(todos) == 0 {
			fmt.Println("There is no todo to delete")
		} else {
			todoId := ""

			for index := 2; index <= len(os.Args)-1; index++ {
				if index == 2 {
					todoId = os.Args[index]
					break
				}
			}

			if todoId == "" {
				fmt.Println("Please input todo id")
				os.Exit(1)
			}

			isExist := false
			for index, todo := range todos {
				if strconv.Itoa(int(todo.Id)) == todoId {
					todosTmp := todos[0:index]
					todosTmp = append(todosTmp, todos[index+1:]...)
					todos = todos[:0]
					todos = append(todos, todosTmp...)
					isExist = true
					break
				}
			}

			if !isExist {
				fmt.Println("Todo not found")
				os.Exit(1)
			}

			json, err := json.Marshal(todos)
			if err != nil {
				fmt.Println("Failed parse struct to json")
			}
			file.Truncate(0)
			if _, err := file.Write(json); err != nil {
				fmt.Println("Failed save todo")
			}
		}
		break
	case "mark-in-progress":
		if len(todos) == 0 {
			fmt.Println("There is no todo to delete")
		} else {
			todoId := ""

			for index := 2; index <= len(os.Args)-1; index++ {
				if index == 2 {
					todoId = os.Args[index]
					break
				}
			}

			if todoId == "" {
				fmt.Println("Please input todo id")
				os.Exit(1)
			}

			isExist := false
			for index, todo := range todos {
				if strconv.Itoa(int(todo.Id)) == todoId {
					todos[index].Status = "in-progress"
					todos[index].UpdatedAt = time.Now().Format(time.DateTime)
					isExist = true
					break
				}
			}

			if !isExist {
				fmt.Println("Todo not found")
				os.Exit(1)
			}

			json, err := json.Marshal(todos)
			if err != nil {
				fmt.Println("Failed parse struct to json")
			}
			file.Truncate(0)
			if _, err := file.Write(json); err != nil {
				fmt.Println("Failed save todo")
			}
		}
		break
	case "mark-done":
		if len(todos) == 0 {
			fmt.Println("There is no todo to delete")
		} else {
			todoId := ""

			for index := 2; index <= len(os.Args)-1; index++ {
				if index == 2 {
					todoId = os.Args[index]
					break
				}
			}

			if todoId == "" {
				fmt.Println("Please input todo id")
				os.Exit(1)
			}

			isExist := false
			for index, todo := range todos {
				if strconv.Itoa(int(todo.Id)) == todoId {
					todos[index].Status = "done"
					todos[index].UpdatedAt = time.Now().Format(time.DateTime)
					isExist = true
					break
				}
			}

			if !isExist {
				fmt.Println("Todo not found")
				os.Exit(1)
			}

			json, err := json.Marshal(todos)
			if err != nil {
				fmt.Println("Failed parse struct to json")
			}
			file.Truncate(0)
			if _, err := file.Write(json); err != nil {
				fmt.Println("Failed save todo")
			}
		}
		break
	case "list":
		if len(todos) == 0 {
			fmt.Println("There is no todo to list")
		} else {
			status := ""
			for index := 2; index <= len(os.Args)-1; index++ {
				if index == 2 {
					status = os.Args[index]
					break
				}
			}

			fmt.Println("|Id|Todo|Status")
			for _, todo := range todos {
				if status == "" {
					fmt.Print(fmt.Sprintf("|%v|%v|%v\n", todo.Id, todo.Description, todo.Status))
				} else {
					if status == "done" && todo.Status == "done" {
						fmt.Print(fmt.Sprintf("|%v|%v|%v\n", todo.Id, todo.Description, todo.Status))
					}

					if status == "todo" && todo.Status == "todo" {
						fmt.Print(fmt.Sprintf("|%v|%v|%v\n", todo.Id, todo.Description, todo.Status))
					}

					if status == "in-progress" && todo.Status == "in-progress" {
						fmt.Print(fmt.Sprintf("|%v|%v|%v\n", todo.Id, todo.Description, todo.Status))
					}
				}
			}
		}
		break
	default:
		fmt.Printf("Unknow command %s\n", os.Args[1])
		fmt.Println("Available command: ")
		fmt.Println("1. add [todo]")
		fmt.Println("2. update [todo id] [new todo]")
		fmt.Println("3. delete [todo id]")
		fmt.Println("4. mark-in-progress [todo id]")
		fmt.Println("5. mark-done [todo id]")
		fmt.Println("6. list [done|todo|in-progress]")
		break
	}
}
