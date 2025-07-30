package internal

import (
	"cli-task-tracker/task"
	"fmt"
	"os"
)

func HandleCommand(args []string) {
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
		newTask := task.Task{}
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
