package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		// Step 4: implement ping command
		if strings.Compare("ping", commandString[0]) == 0 {
			commandOutput := execPINGCommand(commandString[1])
			fmt.Println(commandOutput)
		} else {
			fmt.Println("Command not supported by the CLI")
		}

		fmt.Print("=>")
	}

}

func execPINGCommand(cmdParam string) string {
	// Step 5: Execute OS command ping
	c, err := exec.Command("ping", "-c", "3", cmdParam).CombinedOutput()
	if err != nil {
		return "Error! Command execution failed" + err.Error()
	}

	return string(c)
}
