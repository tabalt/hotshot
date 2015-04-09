package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fmt.Println("start server")

	startServer()

	fmt.Println("end server")
}

func startServer() {
	service := ":4732"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("listen success: ", tcpAddr, "\n\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	fmt.Println("start  handleClient")

	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout

	defer conn.Close() // close connection before exit
	for {
		// set maxium request length to 128KB to prevent flood attack

		//通信协议解包
		magicByte := make([]byte, 4)
		typeByte := make([]byte, 4)
		sizeByte := make([]byte, 4)

		_, _ = conn.Read(magicByte)
		_, _ = conn.Read(typeByte)
		_, _ = conn.Read(sizeByte)

		var typeValue, sizeValue int32

		magicValue := string(magicByte)

		buf := bytes.NewBuffer(typeByte)
		binary.Read(buf, binary.BigEndian, &typeValue)
		buf = bytes.NewBuffer(sizeByte)
		binary.Read(buf, binary.BigEndian, &sizeValue)

		dataByte := make([]byte, sizeValue)
		_, _ = conn.Read(dataByte)

		dataValue := string(dataByte)

		fmt.Println("magic :", magicValue, "\ttype :", typeValue, "\tsize :", sizeValue, "\tdata :", dataValue)

		//os.Exit(0)

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
