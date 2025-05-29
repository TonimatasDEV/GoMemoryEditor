package system

import "syscall"

const (
	PROCESS_VM_WRITE          = 0x0020
	PROCESS_VM_OPERATION      = 0x0008
	PROCESS_VM_READ           = 0x0010
	PROCESS_QUERY_INFORMATION = syscall.PROCESS_QUERY_INFORMATION
	MEM_COMMIT                = 0x1000
	PAGE_READWRITE            = syscall.PAGE_READWRITE
)

var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	ProcEnumProcesses      = modPsapi.NewProc("EnumProcesses")
	modPsapi               = syscall.NewLazyDLL("psapi.dll")
	ProcGetModuleBaseNameW = modPsapi.NewProc("GetModuleBaseNameW")
	ProcEnumProcessModules = modPsapi.NewProc("EnumProcessModules")
	ReadProcessMemory      = kernel32.NewProc("ReadProcessMemory")
	WriteProcessMemory     = kernel32.NewProc("WriteProcessMemory")
	VirtualQueryEx         = kernel32.NewProc("VirtualQueryEx")
	GlobalMemoryStatusEx   = kernel32.NewProc("GlobalMemoryStatusEx")
	Processes              = []string{"System", "Idle", "smss.exe", "csrss.exe", "wininit.exe", "winlogon.exe", "services.exe", "lsass.exe", "svchost.exe",
		"sihost.exe", "Explorer.EXE", "ShellHost.exe", "taskhostw.exe", "SearchHost.exe", "StartMenuExperienceHost.exe", "RuntimeBroker.exe",
		"msedgewebview2.exe", "powershell.exe", "full-line-inference.exe", "conhost.exe", "cef_server.exe", "fsnotifier.exe", "GameBarPresenceWriter.exe",
		"GameOverlayUI.exe", "XboxGameBarSpotify.exe", "XboxPcAppFT.exe", "ApplicationFrameHost.exe", "OpenConsole.exe", "backgroundTaskHost.exe",
		"TextInputHost.exe", "LockApp.exe", "SystemSettings.exe", "goland64.exe", "ShellExperienceHost.exe"}
)
