package main

import (
	"bufio"
	"cli-task-tracker/internal"
	"fmt"
	"os"
	"strings"
)

func main() {
	internal.Help() // Display the help message when the program starts
	reader := bufio.NewReader(os.Stdin) // Create a new reader to read input from the console

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n') // Read input until a newline character

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input) // Trim whitespace and \n from the input

		args := strings.Fields(input)
		if len(args) == 0 {
			fmt.Println("No command entered. Type 'help' for a list of commands.")
			continue // If no command is entered, prompt the user again
		}

		internal.HandleCommand(args) // Call the function to handle the command
	}
}