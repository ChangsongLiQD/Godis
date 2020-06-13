package main

import (
	"golang.org/x/sys/unix"
	"math/rand"
	"time"
)

var server *redisServer

type redisServer struct {
	pid int
	//commands
}

func main(){
	rand.Seed(time.Now().Unix())
	initRedisServer()
}

func initRedisServer(){
	server.pid = unix.Getpid()
}
