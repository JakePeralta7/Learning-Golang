package main

import (
	"fmt"
	"syscall"
)

var (
	kernel32            = syscall.MustLoadDLL("kernel32.dll")
	isDebuggerPresent   = kernel32.MustFindProc("IsDebuggerPresent")
)

func IsDebugging() bool {
	debugging, _, _ := isDebuggerPresent.Call()
	return debugging != 0
}

func main() {
	if IsDebugging() {
		fmt.Println("The program is running in debug mode.")
	} else {
		fmt.Println("The program is not running in debug mode.")
	}
}
