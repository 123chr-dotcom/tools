package main

import (
	"fmt"
	"os/exec"
)

func TerminatePort(port int) error {
	// 获取端口信息
	ports := ScanPorts()
	var pid int = -1

	// 查找指定端口的PID
	for _, p := range ports {
		if p.Port == port {
			pid = p.PID
			break
		}
	}

	if pid == -1 {
		return fmt.Errorf("port %d not found", port)
	}

	// 终止进程
	cmd := exec.Command("taskkill", "/F", "/PID", fmt.Sprintf("%d", pid))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to terminate process %d: %v", pid, err)
	}

	return nil
}
