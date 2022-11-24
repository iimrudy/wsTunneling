package client

import (
	"fmt"
	"net"

	"golang.org/x/net/websocket"
)

var sshSocket = ""
var serverSocket = ""
var pathW = ""

func Serve(sshIp string, sshPort string, ip string, port string, path string) {

	sshSocket = sshIp + ":" + sshPort
	serverSocket = ip + ":" + port
	pathW = path
	// create tcp server
	tcpServer, err := net.Listen("tcp", sshSocket)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tcpServer.Close()
	for {
		// accept connection
		tcpConn, err := tcpServer.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		// handle connection
		go handleConnection(tcpConn)
	}
}

func handleConnection(tcpConn net.Conn) {
	defer tcpConn.Close()
	ws, err := websocket.Dial(fmt.Sprintf("ws://%s/%s", serverSocket, pathW), "", fmt.Sprintf("http://%s/", serverSocket))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	go func() {
		for {
			bytes := make([]byte, 1024)
			size, err := ws.Read(bytes)
			if err != nil {
				fmt.Println(err)
				ws.Close()
				return
			}
			_, err = tcpConn.Write(bytes[:size])
			if err != nil {
				fmt.Println(err)
				ws.Close()
				return
			}
		}
	}()

	for {
		bytes := make([]byte, 1024)
		size, err := tcpConn.Read(bytes)
		if err != nil {
			fmt.Println(err)
			ws.Close()
			return
		}
		_, err = ws.Write(bytes[:size])
		if err != nil {
			fmt.Println(err)
			ws.Close()
			return
		}
	}

}
