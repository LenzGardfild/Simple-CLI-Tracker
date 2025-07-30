package internal

import (
	"cli-task-tracker/task"
	"fmt"
	"os"
	"strings"
)

func HandleCommand(args []string, filename string) {
	tasks, _ := task.LoadTasks(filename) // Load tasks from the file

	switch args[0] {
	case "exit":
		fmt.Println("Exiting the program.")
		os.Exit(0) // Exit the program
	case "help":
		Help() // Display the help message
	case "add":
		if len(args) < 2 {
			fmt.Println("Please provide a task to add.")
			return
		}
		text := strings.Join(args[1:], " ")
		maxID := 0
		for _, task := range tasks {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
		newTask := task.Task{ID: maxID + 1, Text: text, Status: "pending"}
		tasks = append(tasks, newTask) // Add the new task to the list
		task.SaveTasks(filename, tasks) // Save the updated task list to the file
		fmt.Printf("Task added: %s\n", text)
	case "delete":
		if len(args) < 2 {
			fmt.Println("Please provide the ID of the task to delete.")
			return
		}
		id := args[1]
		for i, t := range tasks {
			if fmt.Sprint(t.ID) == id {
				tasks = append(tasks[:i], tasks[i+1:]...) // Remove the task from the list
				task.SaveTasks(filename, tasks) // Save the updated task list to the file
				fmt.Printf("Task with ID %s deleted.\n", id)
				return
			}
		}
		fmt.Printf("Task with ID %s not found.\n", id)
	case "list":
		if len(tasks) == 0 {
			fmt.Println("No tasks available.")
			return
		}
		fmt.Println("Tasks:")
		for _, t := range tasks {
			fmt.Printf("ID: %d, Task: %s, Status: %s\n", t.ID, t.Text, t.Status)
		}
	case "update":
		if len(args) < 3 {
			fmt.Println("Please provide the ID of the task to update and the new task text.")
			return
		}
		id := args[1]
		newText := strings.Join(args[2:], " ")
		for i, t := range tasks {
			if fmt.Sprint(t.ID) == id {
				tasks[i].Text = newText // Update the task text
				task.SaveTasks(filename, tasks) // Save the updated task list to the file
				fmt.Printf("Task with ID %s updated to: %s\n", id, newText)
				return
			}
		}
		fmt.Printf("Task with ID %s not found.\n", id)
	case "in-progress":
		if len(args) < 2 {
			fmt.Println("Please provide the ID of the task to mark as in progress.")
			return
		}
		id := args[1]
		for i, t := range tasks {
			if fmt.Sprint(t.ID) == id {
				tasks[i].Status = "in-progress" // Update the task status to in-progress
				task.SaveTasks(filename, tasks) // Save the updated task list to the file
				fmt.Printf("Task with ID %s marked as in progress.\n", id)
				return
			}
		}
		fmt.Printf("Task with ID %s not found.\n", id)
	case "done":
		if len(args) < 2 {
			fmt.Println("Please provide the ID of the task to mark as done.")
			return
		}
		id := args[1]
		for i, t := range tasks {
			if fmt.Sprint(t.ID) == id {
				tasks[i].Status = "done" // Update the task status to done
				task.SaveTasks(filename, tasks) // Save the updated task list to the file
				fmt.Printf("Task with ID %s marked as done.\n", id)
				return
			}
		}
		fmt.Printf("Task with ID %s not found.\n", id)
	default:
		fmt.Println("Unknown command. Use 'help' to see available commands.")
	}
}


func Help() {
	fmt.Print(
`

Simple CLI Task Tracker

 _____________________________________________________
|Use [add] [argument] to add a task                   |
|Use [list] to list all tasks                         |
|Use [update] [ID] [new task] to update a task.       |
|Use [delete] [ID] to delete a task                   |
|Use [in-progress] [ID] to mark a task as in progress |
|Use [done] [ID] to mark a task as done               |
|Use [help] to display this help message              |
|Use [exit] to exit the program                       |
|_____________________________________________________|


`)
}
