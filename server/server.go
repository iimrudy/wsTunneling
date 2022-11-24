package server

import (
	"fmt"
	"net"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

var sshSocket = ""

func Serve(sshIp string, sshPort string, ip string, port string, path string) {

	sshSocket = sshIp + ":" + sshPort
	serverSocket := ip + ":" + port
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/"+path, hello)
	e.Logger.Fatal(e.Start(serverSocket))
}

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		sshConn, err := net.Dial("tcp", sshSocket)
		defer sshConn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			for {
				bytes := make([]byte, 1024)
				size, err := ws.Read(bytes)
				if err != nil {
					fmt.Println(err)
					ws.Close()
					return
				}
				_, err = sshConn.Write(bytes[:size])
				if err != nil {
					fmt.Println(err)
					ws.Close()
					return
				}
			}
		}()

		for {
			bytes := make([]byte, 1024)
			size, err := sshConn.Read(bytes)
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
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
