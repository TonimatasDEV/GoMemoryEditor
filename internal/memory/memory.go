package memory

import (
	"encoding/binary"
	"fmt"
	"github.com/TonimatasDEV/GoMemoryEditor/internal"
	"github.com/TonimatasDEV/GoMemoryEditor/internal/system"
	"strconv"
	"syscall"
	"unsafe"
)

func GetMemoryInfo(address uintptr) (uintptr, BasicMemoryInfo) {
	var memoryInfo BasicMemoryInfo

	ret, _, _ := system.VirtualQueryEx.Call(
		uintptr(internal.SelectedProcess),
		address,
		uintptr(unsafe.Pointer(&memoryInfo)),
		unsafe.Sizeof(memoryInfo),
	)

	return ret, memoryInfo
}

func IsMemoryModifiable(info BasicMemoryInfo) bool {
	return info.State == system.MEM_COMMIT && info.Protect == system.PAGE_READWRITE
}

func ReadMemory(hProcess syscall.Handle, address uintptr, size uintptr) (uintptr, []byte, uintptr) {
	region := make([]byte, size)
	var bytesRead uintptr

	ret, _, _ := system.ReadProcessMemory.Call(
		uintptr(hProcess),
		address,
		uintptr(unsafe.Pointer(&region[0])),
		size,
		uintptr(unsafe.Pointer(&bytesRead)),
	)

	return ret, region, bytesRead
}

func WriteMemory(address uintptr, value int32) (bool, uintptr) {
	buffer := (*[4]byte)(unsafe.Pointer(&value))
	var bytesWritten uintptr

	ret, _, _ := system.WriteProcessMemory.Call(
		uintptr(internal.SelectedProcess),
		address,
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(4),
		uintptr(unsafe.Pointer(&bytesWritten)),
	)

	if ret == 0 {
		return false, 0
	} else {
		return true, bytesWritten
	}
}

func FilterByNewValue(newValue int32) {
	var filtered []uintptr

	for _, address := range internal.FoundAddresses {
		ret, buffer, bytesRead := ReadMemory(internal.SelectedProcess, address, 4)

		if ret == 0 || bytesRead != 4 {
			continue
		}

		currentValue := int32(binary.LittleEndian.Uint32(buffer[:]))
		if currentValue == newValue {
			fmt.Printf("Found changed value at address: 0x%X - %d\n", address, currentValue)
			filtered = append(filtered, address)
		}
	}

	internal.FoundAddresses = filtered
}

func ConvStrToUintptr(str string) (uintptr, error) {
	clean := str
	if len(str) >= 2 && str[:2] == "0x" {
		clean = str[2:]
	}

	addr, err := strconv.ParseUint(clean, 16, 64)
	if err != nil {
		return 0, err
	}

	return uintptr(addr), nil
}

func AddValueWithTargetFromRegion(memoryInfo BasicMemoryInfo, targetValue int32) {
	_, region, bytesRead := ReadMemory(internal.SelectedProcess, memoryInfo.BaseAddress, memoryInfo.RegionSize)

	for i := 0; i < int(bytesRead)-4; i++ {
		value := int32(binary.LittleEndian.Uint32(region[i : i+4]))
		if value == targetValue {
			internal.FoundAddresses = append(internal.FoundAddresses, memoryInfo.BaseAddress+uintptr(i))
			fmt.Printf("Found value at address: 0x%X\n", memoryInfo.BaseAddress+uintptr(i))
		}
	}
}
