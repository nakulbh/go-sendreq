package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	host, path, method string
	port               int
)

func main() {
	flag.StringVar(&method, "method", "GET", "HTTP method to use")
	flag.StringVar(&host, "host", "localhost", "host to connect to")
	flag.IntVar(&port, "port", 8080, "port to connect to ")
	flag.StringVar(&path, "path", "/", "path to request")
	flag.Parse()

	ip, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, ip)
	if err != nil {
		panic(err)
	}

	log.Printf("connected to %s (@ %s )", host, conn.RemoteAddr())
}
