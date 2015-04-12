package server

import (
	"../packet"
	"fmt"
	"github.com/tabalt/appgo/error"
	"net"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {

	return &Server{Port: port}
}

func (this *Server) Start() {

	service := ":" + this.Port

	tcpAddress, err := net.ResolveTCPAddr("tcp4", service)
	error.CheckError(err)

	listener, err := net.ListenTCP("tcp", tcpAddress)
	error.CheckError(err)

	fmt.Println("listen success: ", tcpAddress, "\n\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go this.handleConnection(conn)
	}
}

func (this *Server) handleConnection(conn net.Conn) {

	fmt.Println("start handleConnection")

	// close connection before exit
	//defer conn.Close()

	for {

		pack := packet.BuildPacketFromConnection(conn)

		fmt.Println("magic :", pack.Magic, "\ttype :", pack.Type, "\tsize :", pack.Size, "\tdata :", pack.Data)

		//os.Exit(0)

	}
}
