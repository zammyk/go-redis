package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		resp := NewResp(conn)
		writer := NewWriter(conn)

		value, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}

		_ = value
		writer.Write(Value{typ: "string", str: "OK"})
	}

	defer conn.Close()
}
