# Golang tcp reverse shell
## How does it work?
* The Go code establishes a reverse shell to a Windows machine.
* Commands are executed on the local machine using the Windows cmd.exe shell.
* The console window is hidden when the reverse shell is executed.

## Build the Go shell into a Windows executable
`go build -o goshell.exe goshell.go`

## How to setup a listener on the attacker machine
`nc -nlvp 4444`
