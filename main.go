package main

import (
	"flag"
)

type Player struct {
	Name string
}

var port string
var serverMode bool
var serverAddress string
var name string

// set flags
func init() {
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.BoolVar(&serverMode, "server", false, "run in server mode... this is also a crappy description")
	flag.BoolVar(&serverMode, "s", false, "run in server mode to serve matches. this is a crappy description")
	flag.StringVar(&serverAddress, "a", "127.0.0.1", "battleterm server ip address you wish to connect to")
	flag.StringVar(&serverAddress, "addr", "127.0.0.1", "battleterm server ip address you wish to connect to")
	flag.StringVar(&name, "name", "player", "player's name")
	flag.StringVar(&name, "n", "player", "player's name")
}

func main() {
	flag.Parse()
	port = ":" + port
	if serverMode {
		beServer()
	} else {
		beClient()
	}	
}
