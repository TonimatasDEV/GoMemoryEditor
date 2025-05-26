package actions

import (
	"fmt"
	"github.com/TonimatasDEV/GoMemoryEditor/internal"
	"github.com/TonimatasDEV/GoMemoryEditor/internal/memory"
	"github.com/TonimatasDEV/GoMemoryEditor/internal/process"
	"os"
	"strconv"
)

func Processes(_ []string) {
	for pid, name := range process.GetProcessMapWithNames() {
		fmt.Printf("PID: %-6d Name: %s\n", pid, name)
	}
}

func SelectProcess(args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments.")
		return
	}

	id, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("Invalid arguments.")
		return
	}

	hProcess, err := process.OpenProcess(uint32(id))

	if err != nil {
		fmt.Println("Failed to open process:", err)
		return
	}

	internal.SelectedProcess = hProcess
	fmt.Println("Selected process:", internal.SelectedProcess)
}

func Exit(_ []string) {
	os.Exit(0)
}

func Search(args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments.")
		return
	}

	targetValue, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid arguments")
		return
	}

	internal.FoundAddresses = make([]uintptr, 0)

	fmt.Println("Searching...")

	var address uintptr

	for address = 0x00000000; address < 0x7FFFFFFF; {
		ret, memoryInfo := memory.GetMemoryInfo(address)

		if ret == 0 {
			break
		}

		if memory.IsMemoryModifiable(memoryInfo) {
			memory.AddValueWithTargetFromRegion(memoryInfo, int32(targetValue))
		}

		address = memoryInfo.BaseAddress + memoryInfo.RegionSize
	}
}

func Research(args []string) {
	if len(args) != 1 {
		fmt.Println("Incorrect number of arguments.")
		return
	}

	value, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid arguments")
		return
	}

	fmt.Println("Researching...")

	memory.FilterByNewValue(int32(value))
}

func ModifyAddressValue(args []string) {
	if len(args) != 2 {
		fmt.Println("Invalid arguments.")
		return
	}

	address, err := memory.ConvStrToUintptr(args[0])
	newValue, err1 := strconv.Atoi(args[1])

	if err != nil || err1 != nil {
		fmt.Println("Invalid arguments.")
		return
	}

	ok, bytesWritten := memory.WriteMemory(address, int32(newValue))

	if ok {
		fmt.Printf("Successfully wrote %d to address 0x%X (%d bytes written)\n", newValue, address, bytesWritten)
	} else {
		fmt.Println("Failed to write memory.")
	}
}

func Help(_ []string) {
	help := "Commands:\n"

	for _, command := range commands {
		help += command.Print()
	}

	fmt.Print(help)
}
