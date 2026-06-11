package cmd

import "fmt"

func Help() {
	commandList := map[string]string{
		"add":    "Add event - Syntax: add \"name event\" \"date\" \"priority\"",
		"list":   "Show all events - Syntax: list",
		"remove": "Delete event - Syntax: remove \"id\"",
		"update": "Update event - Syntax: update \"id\" \"name event\" \"date\" \"priority\"",
		"remind": "Set event reminder - Syntax: remind \"id\" \"reminder message\" \"date\"",
		"help":   "Show list of commands - Syntax: help",
		"exit":   "Exit program - Syntax: exit",
	}

	fmt.Println("\nAvailable commands:")
	fmt.Println("-------------------")
	for command, desc := range commandList {
		fmt.Printf("  %-8s - %s\n", command, desc)
	}
}
