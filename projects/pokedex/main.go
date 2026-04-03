package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/eallender/pokedex/internal/repl"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	repl.Init()

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input := repl.CleanInput(scanner.Text())
			command, ok := repl.CliCommands[input[0]]
			if ok {
				err := command.Callback()
				if err != nil {
					fmt.Printf("An error occurred: %s\n", err.Error())
				}
			} else {
				fmt.Print("Unknown command\n")
			}
		}
	}

}
