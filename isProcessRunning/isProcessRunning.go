package main

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func isProcessRunning(processName string) (bool, error) {
	processes, err := process.Processes()
	if err != nil {
		return false, err
	}

	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}
		if name == processName {
			return true, nil
		}
	}

	return false, nil
}

func main() {
	processList := []string{
		"notepad.exe",
		"explorer.exe",
		"chrome.exe"} // replace with your list of process names

	for _, processName := range processList {
		isRunning, err := isProcessRunning(processName)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		if isRunning {
			fmt.Printf("The process %s is running.\n", processName)
		} else {
			fmt.Printf("The process %s is not running.\n", processName)
		}
	}
}
