package main

import (
	"fmt"
	"os"
)

var server = &Server{}

const AppVersion = "Godis 0.1"

func InitServer() {
	server.Pid = os.Getpid()
	server.Port = 6379

}

func populateCommandTable() {

}

func SmoothExit() {
	fmt.Println("Handle finish")
	fmt.Println("bye")
}
