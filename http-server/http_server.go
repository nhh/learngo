package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func log(string string) {
	io.WriteString(os.Stdout, string+"\n")
}

func Serve() {
	ln, serverError := net.Listen("tcp", ":80")
	if serverError != nil {
		log("Err: Listening in port 80")
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log("Err: Accepting tcp connection")
		}

		go func(connection net.Conn) {
			defer connection.Close()

			buf := make([]byte, 8192)

			for {
				_, connReadErr := connection.Read(buf)

				if connReadErr != nil && connReadErr != io.EOF {
					log(connReadErr.Error())
				}

				if connReadErr == io.EOF {
					break
				}
			}

			resp := Handle(string(buf))

			_, connErro := connection.Write([]byte(resp))
			if connErro != nil {
				panic(connErro)
			}

			log("Served request")
		}(conn)
	}

}

func Handle(request string) string {
	log(request)

	rq := strings.Split(request, "\n")

	if len(rq) == 0 {
		panic("Could not parse http request")
	}

	headline := strings.Split(rq[0], " ")

	method := headline[0]
	//path := headline[1]
	//version := headline[2]

	responseCode := 0

	switch method {
	case "GET":
		responseCode = 200
	case "POST":
		responseCode = 201
	default:
		log(fmt.Sprintf("METHOD: %s", method))
		panic("Unsupported http method")
	}

	return fmt.Sprintf("HTTP/1.0 %d OK\n\r", responseCode)
}

func main() {
	Serve()
}
