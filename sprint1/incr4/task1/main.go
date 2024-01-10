package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (naddr *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", naddr.Host, naddr.Port)
}

func (naddr *NetAddress) Set(flagValue string) error {
	var sPort string
	var err error
	naddr.Host, sPort, _ = strings.Cut(flagValue, ":")
	naddr.Port, err = strconv.Atoi(sPort)
	return err
}

func main() {
	addr := new(NetAddress)
	_ = flag.Value(addr)
	flag.Var(addr, "addr", "Net address host:port")
	flag.Parse()
	fmt.Println(addr.Host)
	fmt.Println(addr.Port)
}
