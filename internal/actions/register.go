package actions

import (
	"fmt"
	"strings"
)

var commands = make(map[string]func([]string))

func Init() {
	commands["select"] = SelectProcess
	commands["processes"] = Processes
	commands["search"] = Search
	commands["research"] = Research
	commands["modify"] = ModifyAddressValue
	commands["exit"] = Exit
}

func Run(input string) {
	parts := strings.Fields(input)
	cmd, args := parts[0], parts[1:]

	if commands[cmd] != nil {
		commands[cmd](args)
		return
	}

	fmt.Println("Unknown command:", cmd)
}
