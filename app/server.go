package app

import (
	"Godis/core/cmd"
	"fmt"
)

type Server struct {
	Pid      int
	Port     int
	Clients  int64
	Commands map[string]*cmd.Command
}

func InitServer() {

}

func SmoothExit() {
	fmt.Println("Handle finish")
	fmt.Println("bye")
}
