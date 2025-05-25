package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunController() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// 显示端口列表
		ports := ScanPorts()
		DisplayPorts(ports)

		// 获取用户输入
		fmt.Print("Enter port number to terminate (or 'exit' to quit): ")
		scanner.Scan()
		input := scanner.Text()

		// 处理退出命令
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		// 尝试解析端口号
		port, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid port number.")
			continue
		}

		// 检查端口是否存在
		if !PortExists(port, ports) {
			fmt.Printf("Port %d does not exist. Please try again.\n", port)
			continue
		}

		// 终止端口
		if err := TerminatePort(port); err != nil {
			fmt.Printf("Error terminating port %d: %v\n", port, err)
		} else {
			fmt.Printf("Successfully terminated port %d\n", port)
		}
	}
}

func PortExists(port int, ports []PortInfo) bool {
	for _, p := range ports {
		if p.Port == port {
			return true
		}
	}
	return false
}
