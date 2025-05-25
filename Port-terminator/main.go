package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] != "start" {
		fmt.Println("Usage: port_terminator start")
		os.Exit(1)
	}



	ShowWelcome()

	// 主控制循环
	RunController()
}
