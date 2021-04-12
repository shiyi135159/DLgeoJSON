package tool

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	_SetConsoleTitle uintptr
)

func init() {
	kernel32, loadErr := syscall.LoadLibrary("kernel32.dll")
	if loadErr != nil {
		fmt.Println("loadErr", loadErr)
	}
	defer syscall.FreeLibrary(kernel32)
	_SetConsoleTitle, _ = syscall.GetProcAddress(kernel32, "SetConsoleTitleW")
}

// SetConsoleTitle set the console title.
func SetConsoleTitle(title string) int {
	ret, _, callErr := syscall.Syscall(_SetConsoleTitle, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	if callErr != 0 {
		fmt.Println("callErr", callErr)
	}
	return int(ret)
}
