package scanlib

//this file is used for manage process by name id or etc.

import (
	"fmt"
	"os/exec"
	"winportkill/scanlib/types"
	"strings"
)

func GetProcessName(pid string) string {
	cmd := exec.Command("wmic", "process", "where", fmt.Sprintf("ProcessId=%s", pid), "get", "Name")
	output, err := cmd.Output()
	if err != nil {
		return "Unknown"
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line = strings.TrimSpace(line); line != "" && line != "Name" {
			return line
		}
	}
	return "Unknown"
}

func KillProcess(pid string) bool {
	cmd := exec.Command("taskkill", "/PID", pid, "/F")
	return cmd.Run() == nil
}

func KillAllProcesses(ports []types.ProcessInfo) bool {
	success := true
	for _, p := range ports {
		if KillProcess(p.PID) {
			fmt.Printf("Process with PID %s (%s) terminated successfully.\n", p.PID, p.ProcessName)
		} else {
			fmt.Printf("Failed to terminate process with PID %s (%s).\n", p.PID, p.ProcessName)
			success = false
		}
	}
	return success
}

func KillByName(name string) bool {
	if !strings.HasSuffix(strings.ToLower(name), ".exe") {
		name = name + ".exe"
	}
	cmd := exec.Command("taskkill", "/IM", name, "/F")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to terminate processes with name %s: %v\n", name, err)
		return false
	}
	fmt.Printf("All processes with name %s terminated successfully.\n", name)
	return true
}
