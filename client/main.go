package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("Connected to %v!\n", conn.RemoteAddr())

	go func(conn net.Conn) {
		r := bufio.NewReader(conn)
		for {
			msg, _, err := r.ReadLine()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Server said: %v\n", string(msg))
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
