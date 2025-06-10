package main

import (
	"bufio"
	"fmt"
	"os"
	"winportkill/clicore"
	"winportkill/scanlib"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(clicore.GetAsciiArt())
	fmt.Println("Port Kill CLI - Enter port number to scan (or 'q' to quit):")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		port := strings.TrimSpace(scanner.Text())

		if strings.ToLower(port) == "q" {
			fmt.Println("Exiting...")
			break
		}

		if _, err := strconv.Atoi(port); err != nil {
			fmt.Println("Please enter a valid port number or 'q' to quit.")
			continue
		}

		fmt.Println("Scanning Port ...")
		ports := scanlib.ListOpenPorts(port)
		scanlib.PrintPorts(port, ports)
		if len(ports) == 0 {
			continue
		}

		fmt.Println("\nOptions: Enter PID or process name to kill, 'all' to kill all, 'skip' to continue, or process name (e.g., node):")
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "skip" {
			continue
		}

		if strings.ToLower(input) == "all" {
			if scanlib.KillAllProcesses(ports) {
				fmt.Println("All processes for port", port, "terminated successfully.")
			} else {
				fmt.Println("Some processes could not be terminated.")
			}
			continue
		}

		for _, p := range ports {
			if p.PID == input {
				if scanlib.KillProcess(input) {
					fmt.Printf("Process with PID %s (%s) terminated successfully.\n", input, p.ProcessName)
				} else {
					fmt.Printf("Failed to terminate process with PID %s (%s).\n", input, p.ProcessName)
				}
				continue
			}
		}

		scanlib.KillByName(input)
	}
}
