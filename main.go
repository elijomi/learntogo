package main

import (
	"fmt"
	tasktracker "learntogo/tasktracker"
	"os"
	"strconv"
)

var taskList *tasktracker.TaskList // todo swap out with persist to json file

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tasktracker <command> [options]")
		return
	}

	taskList = tasktracker.NewTaskList("Daily Tasks")

	// todo add a continuous prompt

	switch os.Args[1] {
	case "list":
		list()
	case "add":
		if len(os.Args) < 6 || os.Args[2] != "--id" {
			fmt.Println("Usage: tasktracker add --title <string> --priority <number>")
			return
		}
		title := os.Args[3]
		priority, err := strconv.Atoi(os.Args[4])
		if err != nil {
			fmt.Println("Invalid priority:", os.Args[4])
			return
		}

		add(title, priority)
	case "done":
		if len(os.Args) < 4 || os.Args[2] != "--id" {
			fmt.Println("Usage: tasktracker done --id <number>")
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid id:", os.Args[3])
			return
		}

		done(id)
	case "delete":
		if len(os.Args) < 4 || os.Args[2] != "--id" {
			fmt.Println("Usage: tasktracker delete --id <number>")
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid id:", os.Args[3])
			return
		}

		delete(id)
	}
}

func add(title string, priority int) {
	task := tasktracker.NewTask(title, priority)
	taskList.AddTask(task)
}

func list() {
	fmt.Println(taskList)
}

func done(id int) {
	task := taskList.Find(id)
	if task == nil {
		fmt.Println("No such task!")
		return
	}

	task.Done = true
	fmt.Printf("Task #%d marked as done.\n", id)
}

func delete(id int) {
	task := taskList.Delete(id)
	if task == nil {
		fmt.Println("No such task!")
		return
	}
	fmt.Printf("Deleted task #%d.\n", id)
}
