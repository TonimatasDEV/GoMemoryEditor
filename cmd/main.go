package main

import (
	"bufio"
	"fmt"
	"github.com/TonimatasDEV/GoMemoryEditor/internal/actions"
	"os"
)

func main() {
	actions.Init()

	fmt.Println("Welcome to GoMemoryEditor! Use \"help\" to see our command.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		actions.Run(input)
	}
}
