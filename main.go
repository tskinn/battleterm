package main

import (
	"flag"
)

type Player struct {
	Name string
}

var port string
var serverMode bool

// set flags
func init() {
	flag.StringVar(&port, "port", "8080", "port to listen on")
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.BoolVar(&serverMode, "s", false, "run in server mode to serve matches. this is a crappy description")
	flag.BoolVar(&serverMode, "server", false, "run in server mode... this is also a crappy description")
}


// listen on specified port for players and add connections to channel
// i guess the channel isn't necessary any more because this isn't
// concurrent
func main() {
	flag.Parse()
	port = ":" + port
	if serverMode {
		beServer()
	} else {
		beClient()
	}
	
}
