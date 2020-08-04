package main

import (
	"Godis/app"
	"Godis/core/parser"
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

	fs := parser.GetFlagSet()
	if err := fs.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if fs.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(app.Version)
		os.Exit(0)
	}

	app.InitServer()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	for sig := range c {
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			app.SmoothExit()
			os.Exit(0)
		}
	}
}

func initServer() {

}
