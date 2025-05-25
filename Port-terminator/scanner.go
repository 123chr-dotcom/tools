package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type PortInfo struct {
	Port     int
	Protocol string
	PID      int
	Process  string
}

func ScanPorts() []PortInfo {
	cmd := exec.Command("netstat", "-ano")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	return parseNetstatOutput(string(output))
}

func DisplayPorts(ports []PortInfo) {
	fmt.Println("\nActive Ports:")
	fmt.Println("Port\tProtocol\tPID\tProcess")
	for _, p := range ports {
		fmt.Printf("%d\t%s\t\t%d\t%s\n", p.Port, p.Protocol, p.PID, p.Process)
	}
	fmt.Printf("\nTotal ports detected: %d\n\n", len(ports))
}

func parseNetstatOutput(output string) []PortInfo {
	var ports []PortInfo
	scanner := bufio.NewScanner(strings.NewReader(output))

	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "LISTENING") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}

		// 解析本地地址和端口
		localAddr := fields[1]
		portStr := strings.Split(localAddr, ":")[1]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			continue
		}

		// 获取协议
		protocol := fields[0]

		// 获取PID
		pid, err := strconv.Atoi(fields[4])
		if err != nil {
			continue
		}

		// 获取进程名称
		process := getProcessName(pid)

		ports = append(ports, PortInfo{
			Port:     port,
			Protocol: protocol,
			PID:      pid,
			Process:  process,
		})
	}

	return ports
}

func getProcessName(pid int) string {
	cmd := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/FO", "CSV", "/NH")
	output, err := cmd.Output()
	if err != nil {
		return "Unknown"
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		return "Unknown"
	}

	// 解析CSV格式的输出
	reader := csv.NewReader(strings.NewReader(lines[0]))
	record, err := reader.Read()
	if err != nil || len(record) < 1 {
		return "Unknown"
	}

	// 移除引号
	processName := strings.Trim(record[0], `"`)
	if processName == "" {
		return "Unknown"
	}

	return processName
}
