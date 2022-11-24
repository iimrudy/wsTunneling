package main

import (
	"flag"
	"fmt"
	"sshbypass/server"
	"sshbypass/client"
)

func main() {
	// parse flag mode from args
	mode := flag.String("mode", "", "specify the mode, if server or client mode")
	sshIp := flag.String("ssh-ip", "127.0.0.1", "specify the ip of the ssh server ")
	sshPort := flag.String("ssh-port", "22", "specify the port of the ssh server")
	ip := flag.String("ip", "0.0.0.0", "specify the ip of the server wrapper ")
	port := flag.String("port", "8080", "specify the port of the server wrapper ")
	path := flag.String("path", "ssh", "specify the path ")

	flag.Parse()

	// check if mode is server or client
	if *mode == "server" {
		fmt.Println("server mode")
		server.Serve(*sshIp, *sshPort, *ip, *port, *path)
	} else if *mode == "client" {
		fmt.Println("client mode")
		client.Serve(*sshIp, *sshPort, *ip, *port, *path)
	} else {
		panic(fmt.Errorf("invalid mode"))
	}

}
