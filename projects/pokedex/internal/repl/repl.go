package repl

import (
	"fmt"
	"os"
	"strings"
)

var CliCommands map[string]cliCommand

func Init() {
	CliCommands = map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays the help message",
			Callback:    commandHelp,
		},
	}
}

func CleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\n")
	fmt.Print("Usage:\n\n")

	for key := range CliCommands {
		fmt.Printf("%s, %s\n", CliCommands[key].Name, CliCommands[key].Description)
	}

	return nil
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}
