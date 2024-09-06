package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Println(userInput)
		fmt.Print("=>")
	}
}
