package internal

import "syscall"

var (
	SelectedProcess syscall.Handle
	FoundAddresses  []uintptr
)
