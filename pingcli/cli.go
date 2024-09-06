package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Display Starting Prompt
	fmt.Println("Cloud CLI")
	fmt.Println("------------------")
	fmt.Print("=>")

	// Accept Command Line Input
	// Example:
	// =>exit
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()
		var commandString []string = strings.Split(userInput, " ")

		switch strings.ToLower(commandString[0]) {
		case "exit":
			// implement exit Command
			os.Exit(0)
		case "ping":
			// implement ping command
			if argsLen := len(commandString); argsLen < 2 {
				fmt.Println("Command not supported by the CLI")
			} else {
				count := "3"
				if argsLen > 2 {
					count = commandString[2]
				}

				commandOutput := execPINGCommand(commandString[1], count)

				fmt.Println(commandOutput)
			}
		default:
			fmt.Println("Command not supported by the CLI")
		}

		fmt.Print("=>")
	}

}

func execPINGCommand(cmdParam, count string) string {
	// Execute OS command ping
	c, err := exec.Command("ping", "-c", count, cmdParam).CombinedOutput()
	if err != nil {
		return "Error! Command execution failed: " + err.Error()
	}

	return string(c)
}
