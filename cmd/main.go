package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"

	"github.com/KeyZenDB/KeyZenDB/config"
)

func main() {
	configPath := flag.String("config", "config.toml", "config file")
	flag.Parse()
	_ = config.LoadConfig(*configPath)
	ln, err := net.Listen("tcp", ":6379") // Redis default port
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("TCP Server listening on port 6379...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(conn)

			for {
				command, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Connection closed")
					return
				}
				command = strings.TrimSpace(command)
				fmt.Println("Received:", command)

				// Process command (example: PING)
				if command == "PING" {
					conn.Write([]byte("PONG\n"))
				} else {
					conn.Write([]byte("UNKNOWN COMMAND\n"))
				}
			}
		}(conn)
	}
}
