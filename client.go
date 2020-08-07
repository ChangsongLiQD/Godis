package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var address string

func main() {
	reader := bufio.NewReader(os.Stdin)
	address = fmt.Sprintf("%s:%d", DefaultIp, DefaultPort)
	tcpAddr, err := net.ResolveTCPAddr("tpc4", address)
	clientCheckError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	clientCheckError(err)
	defer conn.Close()

	for {
		printConsole()
		cmd, err := reader.ReadString('\n')
		clientCheckError(err)
		cmd = strings.Trim(cmd, "\n")

		_, err = sendCmdToServer(cmd, conn)
		clientCheckError(err)

		buff := make([]byte, 1024)
		n, err := conn.Read(buff)
		clientCheckError(err)

		if n == 0 {
			printConsole("(nil)")
		} else if err == nil {
			printConsole(string(buff))
		} else {
			printConsole("server response err")
		}
	}
}

func sendCmdToServer(cmd string, conn *net.TCPConn) (n int, err error) {
	return conn.Write([]byte(cmd))
}

func clientCheckError(err error) {
	if err != nil {
		log.Printf("err: %s", err.Error())
		os.Exit(1)
	}
}

func printConsole(infos ...string) {
	fmt.Print(address + "> ")
	if len(infos) > 0 {
		for _, info := range infos {
			fmt.Print(info)
		}
		fmt.Println()
	}
}
