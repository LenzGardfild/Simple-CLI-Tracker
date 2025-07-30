package internal

import "fmt"



func Help() {
	fmt.Println(
` _____________________________________________________
|Use [add] [argument] to add a task                   |
|Use [list] to list all tasks                         |
|Use [list] [ID] to list a specific task              |
|Use [update] [ID] [new task] to update a task.       |
|Use [delete] [ID] to delete a task                   |
|Use [in-progress] [ID] to mark a task as in progress |
|Use [done] [ID] to mark a task as done               |
|Use [help] to display this help message              |
|_____________________________________________________|`)
}
