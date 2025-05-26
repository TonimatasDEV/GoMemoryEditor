package process

import (
	"github.com/TonimatasDEV/GoMemoryEditor/internal/system"
	"strings"
	"syscall"
	"unsafe"
)

func OpenProcess(pid uint32) (syscall.Handle, error) {
	return syscall.OpenProcess(
		system.PROCESS_VM_WRITE|system.PROCESS_VM_OPERATION|system.PROCESS_QUERY_INFORMATION|system.PROCESS_VM_READ,
		false,
		pid,
	)
}

func GetProcessMainModule(hProcess syscall.Handle) (uintptr, uintptr, uint32) {
	var hMod uintptr
	var cbNeeded uint32

	ret, _, _ := system.ProcEnumProcessModules.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&hMod)),
		unsafe.Sizeof(hMod),
		uintptr(unsafe.Pointer(&cbNeeded)),
	)

	return ret, hMod, cbNeeded
}

func GetProcessExeName(hProcess syscall.Handle, hMod uintptr) (uintptr, string) {
	var exeName [260]uint16

	ret, _, _ := system.ProcGetModuleBaseNameW.Call(
		uintptr(hProcess),
		hMod,
		uintptr(unsafe.Pointer(&exeName[0])),
		uintptr(len(exeName)),
	)

	name := syscall.UTF16ToString(exeName[:])

	return ret, name
}

func GetProcessList() (bool, [1024]uint32) {
	var processIDs [1024]uint32
	var bytesReturned uint32

	r1, _, _ := system.ProcEnumProcesses.Call(
		uintptr(unsafe.Pointer(&processIDs[0])),
		uintptr(len(processIDs))*4,
		uintptr(unsafe.Pointer(&bytesReturned)),
	)

	if r1 == 0 {
		return false, processIDs
	}

	return true, processIDs
}

func GetProcessMapWithNames() map[uint32]string {
	var idMap = make(map[uint32]string)

	ok, processIDs := GetProcessList()

	if !ok {
		return nil
	}

	for _, pid := range processIDs {
		name := GetProcessName(pid)

		if name == "" {
			continue
		}

		if isSystemProcess(name) {
			continue
		}

		idMap[pid] = name
	}

	return idMap
}

func GetProcessName(pid uint32) string {
	hProcess, err := OpenProcess(pid)

	if err != nil || hProcess == 0 {
		return ""
	}

	ret, hMod, _ := GetProcessMainModule(hProcess)

	if ret == 0 {
		CloseHandle(hProcess)
		return ""
	}

	ret, name := GetProcessExeName(hProcess, hMod)
	defer CloseHandle(hProcess)

	if ret == 0 {
		return ""
	}

	return name
}

func CloseHandle(h syscall.Handle) {
	_ = syscall.CloseHandle(h)
}

func isSystemProcess(name string) bool {
	for _, p := range system.Processes {
		if strings.EqualFold(name, p) {
			return true
		}
	}
	return false
}
