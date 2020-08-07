package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const (
	AppVersion = "Godis 0.1"

	DefaultIp   = "127.0.0.1"
	DefaultPort = 6666
)

var server = &Server{}

type Server struct {
	Pid            int
	Port           int
	Clients        int64
	Commands       map[string]Command
	PubSubChannels *map[string]*List
	Db             *Database
}

type Client struct {
	Cmd   Command
	Argv  []Object
	Argc  int
	Query string
	Buff  []byte
}

func InitServer() {
	server.Pid = os.Getpid()
	server.Port = DefaultPort
	server.Db = NewDatabase()
	server.populateCommandTable()
}

func RunServer() {
	addr := fmt.Sprintf("%s:%d", DefaultIp, server.Port)
	netListen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Listen tcp failed: %v", err)
	}

	for {
		conn, err := netListen.Accept()
		if err != nil {
			log.Printf("netListen accept error: %v\n", err)
			continue
		}

		go Handle(conn)
	}

	defer netListen.Close()
}

func Handle(conn net.Conn) {
	client := NewClient()
	for {
		err := client.ReadQuery(conn)
		if err != nil {
			return
		}

		// Process query
		err = client.ProcessInput()
		if err != nil {
			log.Println("client.ProcessInput error:", err)
			continue
		}
		// Process command and Response
		client.HandleCmd()
		Response(conn, client)
	}
}

func Response(conn net.Conn, c *Client) {
	var err error
	if len(c.Buff) > 0 {
		_, err = conn.Write(c.Buff)
	} else {
		_, err = conn.Write([]byte("(nil)"))
	}

	if err != nil {
		log.Println("conn write failed: ", err)
	}
}

func (s *Server) populateCommandTable() {
	s.Commands = map[string]Command{
		"get": GetCommand,
		"set": SetCommand,
	}
}

func (s *Server) getCommand(cmd string) (Command, bool) {
	c, exists := s.Commands[cmd]
	return c, exists
}

func NewClient() *Client {
	return &Client{
		Buff: make([]byte, 512),
	}
}

func (cl *Client) ReadQuery(conn net.Conn) error {
	cl.Buff = make([]byte, 512)
	n, err := conn.Read(cl.Buff)
	if err != nil {
		log.Println("conn.Read error:", err)
		_ = conn.Close()
		return err
	}

	cl.Query = string(cl.Buff[:n])
	return nil
}

func (cl *Client) ProcessInput() error {
	cl.Query = strings.TrimRight(cl.Query, "\n")
	inputs := strings.Split(cl.Query, " ")
	cl.Argc = len(inputs)
	if cl.Argc < 2 {
		return errors.New("invalid command: check usage")
	}
	cl.Argv = make([]Object, cl.Argc-1)

	for idx, input := range inputs {

		if idx == 0 {
			if cmd, exists := server.getCommand(input); exists {
				cl.Cmd = cmd
			} else {
				return errors.New("invalid command")
			}
		} else {
			cl.Argv[idx-1] = CreateObject(ObjectTypeString, input)
		}
	}

	return nil
}

func (cl *Client) HandleCmd() {
	cl.Cmd(cl, server)
}

func SmoothExit() {
	fmt.Println("Handle finish")
	fmt.Println("bye")
}
