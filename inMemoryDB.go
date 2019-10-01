package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			continue
		}

		switch strings.ToUpper(fs[0]) {
		case "GET":
			key := fs[1]
			value := data[key]
			fmt.Fprintf(conn, "%\r\n", value)

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE\r\n")
				continue
			}
			key := fs[1]
			value := fs[2]
			data[key] = value

		case "DEL":
			key := fs[1]
			delete(data, key)

		default:
			fmt.Fprintf(conn, "INVALID COMMAND "+fs[0]+"\r\n")
			continue

		}
	}
}
