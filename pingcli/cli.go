package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Step1: Display Starting Prompt
	fmt.Println("Cloud CLI")
	fmt.Println("------------------")
	fmt.Print("=>")
	// Step 2: Accept Command Line Input
	// Example:
	// =>exit
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()

		var commandString []string = strings.Split(userInput, " ")

		// Step 3: implement exit Command
		if strings.Compare("exit", commandString[0]) == 0 {
			os.Exit(0)
		}

		fmt.Print("=>")
	}
}
