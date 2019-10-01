package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
