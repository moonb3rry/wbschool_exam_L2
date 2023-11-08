package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("shell> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Usage: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Usage: kill <PID>")
				continue
			}
			cmd := exec.Command("taskkill", "/PID", args[1], "/F")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		case "ps":
			cmd := exec.Command("tasklist")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		case "nc":
			if len(args) < 4 {
				fmt.Println("Usage: nc <tcp/udp> <host> <port>")
				continue
			}
			protocol, host, port := args[1], args[2], args[3]
			address := net.JoinHostPort(host, port)
			conn, err := net.Dial(protocol, address)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error connecting:", err)
				continue
			}
			defer conn.Close()

			fmt.Printf("Connected to %s (%s)\n", address, protocol)

			go func() {
				_, err := io.Copy(conn, os.Stdin)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error writing to connection:", err)
				}
			}()

			_, err = io.Copy(os.Stdout, conn)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading from connection:", err)
			}
		case "exit":
			os.Exit(0)
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}
}
