package main

import (
	"fmt"
	"strconv"
)

func ShowWelcome() {
	
	fmt.Println(`
  _____           _            _____      _                                 
 |  __ \         | |          |  __ \    | |                                
 | |__) |__  _ __| |_ ___ _ __| |__) |___| |_ _ __ ___ _ __ ___   ___ _ __ 
 |  ___/ _ \| '__| __/ _ \ '__|  _  // _ \ __| '__/ _ \ '_ ' _ \ / _ \ '__|
 | |  | (_) | |  | ||  __/ |  | | \ \  __/ |_| | |  __/ | | | | |  __/ |   
 |_|   \___/|_|   \__\___|_|  |_|  \_\___|\__|_|  \___|_| |_| |_|\___|_|   
	`)
	fmt.Println("Welcome to Port-terminator!")
	fmt.Println("版权所有 © 2025 zs-yueg")
	fmt.Println("这只是tools中的其中一个工具")
	fmt.Println("更多工具请访问这个github地址: https://github.com/123chr-dotcom/tools")
	fmt.Print("输入端口号以终止该端口: ")
	fmt.Println()
	fmt.Println("输入exit以退出程序")
}

func ShowHelp() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("- Enter port number: Terminate the specified port")
	fmt.Println("- exit: Quit the program")
	fmt.Println()
}

func ValidatePortInput(input string) (int, error) {
	port, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid port number")
	}
	if port < 0 || port > 65535 {
		return 0, fmt.Errorf("port number must be between 0 and 65535")
	}
	return port, nil
}
