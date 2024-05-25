package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	Port   = ":52001" // порт сервера
	MaxLen = 1024     // максимальный размер слайса
)

// handleConn обрабатывает запросы и вычисляет среднее арифметическое.
func handleConn(c net.Conn) {
	defer c.Close()

	for {
		b := make([]byte, MaxLen)
		n, err := c.Read(b)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		var sum int
		for i := 0; i <= n; i++ {
			sum += int(b[i])
		}
		_, err = c.Write([]byte{byte(sum / n)})
		if err != nil {
			fmt.Println("Error writing:", err)
			return
		}
	}
}

// TCPServer запускает сервер и ожидает соединений.
func TCPServer(addr *net.TCPAddr) {
	s, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	defer s.Close()
	for {
		conn, err := s.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err)
			continue
		}
		go handleConn(conn)
	}
}
