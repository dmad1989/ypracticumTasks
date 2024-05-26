package main

import (
	"fmt"
	"net"
)

func UDPClient(addr *net.UDPAddr) {
	s, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Println("connection failed!", err)
		return
	}
	defer s.Close()

	for i := 1; i <= 1000; i++ {
		_, err := fmt.Fprintf(s, "%d", i)
		if err != nil {
			fmt.Println("send data failed!", err)
			return
		}
	}
}

func main() {
	b := make([]byte, 1024)
	addr := &net.UDPAddr{
		Port: 52001,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	go UDPClient(addr)
	var count int
	for {
		_, remoteaddr, err := conn.ReadFromUDP(b)
		if err != nil {
			fmt.Printf("error:  %v\n", err)
			continue
		}
		count++
		if count%100 == 0 {
			fmt.Printf("read from %v, i = %s, count = %d\n", remoteaddr, b, count)
		}
	}
}
