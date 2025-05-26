package main

import (
	"bufio"
	"fmt"
	"github.com/TonimatasDEV/GoMemoryEditor/internal/actions"
	"os"
)

func main() {
	actions.Init()
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
