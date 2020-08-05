package main

type Server struct {
	Pid            int
	Port           int
	Clients        int64
	Commands       map[string]*Command
	PubSubChannels *map[string]*List
}

type Client struct {
}
