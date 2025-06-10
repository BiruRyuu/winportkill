package scanlib

import (
	"fmt"
	"os"
	"os/exec"
	"winportkill/scanlib/types"
	"strings"
	"text/tabwriter"
)

func ListOpenPorts(port string) []types.ProcessInfo {
	var results []types.ProcessInfo
	cmd := exec.Command("netstat", "-ano")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error running netstat: %v\n", err)
		return results
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, ":"+port) && strings.Contains(line, "LISTENING") {
			parts := strings.Fields(line)
			if len(parts) >= 5 {
				results = append(results, types.ProcessInfo{
					Address:     parts[1],
					PID:         parts[4],
					ProcessName: GetProcessName(parts[4]),
				})
			}
		}
	}
	return results
}

func PrintPorts(port string, ports []types.ProcessInfo) {
	if len(ports) == 0 {
		fmt.Printf("No open ports found for port %s.\n", port)
		return
	}

	fmt.Printf("\nOpen ports for %s:\n", port)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)
	_, err := fmt.Fprintln(w, "Address\tPID\tProcess Name\t")
	if err != nil {
		return
	}
	_, err = fmt.Fprintln(w, "-------\t---\t------------\t")
	if err != nil {
		return
	}
	for _, p := range ports {
		fmt.Fprintf(w, "%s\t%s\t%s\t\n", p.Address, p.PID, p.ProcessName)
	}
	w.Flush()
}
