package actions

import (
	"fmt"
	"strings"
)

var commands []CommandInfo

func Init() {
	register("processes", Processes, "Print all processes with they IDs.", "")
	register("select", SelectProcess, "Select the process using IDs.", "[id]")
	register("search", Search, "Print all addresses with X value.", "[value]")
	register("research", Research, "Print all addresses of the last search with X value.", "[value]")
	register("modify", ModifyAddressValue, "Modify X address with Y value.", "[address] [value]")
	register("help", Help, "Print this help message.", "")
	register("exit", Exit, "Exit.", "")
}

func Run(input string) {
	parts := strings.Fields(input)
	cmd, args := parts[0], parts[1:]

	for _, command := range commands {
		if command.Name == cmd {
			command.Function(args)
			return
		}
	}

	fmt.Println("Unknown command:", cmd)
}

func register(id string, function func([]string), description string, arguments string) {
	commands = append(commands, CommandInfo{id, function, description, arguments})
}
