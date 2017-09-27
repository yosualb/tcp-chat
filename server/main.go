package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Connected to %v!\n", conn.RemoteAddr())

	go func(conn net.Conn) {
		r := bufio.NewReader(conn)
		for {
			msg, _, err := r.ReadLine()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Client said: %v\n", string(msg))
		}
	}(conn)

	r := bufio.NewReader(os.Stdin)
	for {
		text, _, err := r.ReadLine()
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(conn, string(text))
	}
}
