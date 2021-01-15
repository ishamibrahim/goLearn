package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprintln(conn, map[string]string{"message": "This is the server"})

		conn.Close()
	}
}
