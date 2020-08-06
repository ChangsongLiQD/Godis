package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	fs := GetFlagSet()
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if fs.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	InitServer()
	go RunServer()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for sig := range c {
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			SmoothExit()
			os.Exit(0)
		}
	}
}
