package memory

import (
	"fmt"
	"github.com/TonimatasDEV/GoMemoryEditor/internal/system"
	"unsafe"
)

type StatusEx struct {
	DwLength                uint32
	DwMemoryLoad            uint32
	UllTotalPhys            uint64
	UllAvailPhys            uint64
	UllTotalPageFile        uint64
	UllAvailPageFile        uint64
	UllTotalVirtual         uint64
	UllAvailVirtual         uint64
	UllAvailExtendedVirtual uint64
}

func GetMaxAddress() uint64 {
	memStatus := StatusEx{}
	memStatus.DwLength = uint32(unsafe.Sizeof(memStatus))

	ret, _, err := system.GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memStatus)))
	if ret == 0 {
		panic(fmt.Sprintf("Error calling GlobalMemoryStatusEx: %v", err))
	}

	return memStatus.UllTotalPhys
}
