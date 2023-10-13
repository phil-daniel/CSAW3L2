package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	//	"net/rpc"
	//	"fmt"
	//	"time"
	//	"net"
)

var nextAddr string

func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.

	listener, _ := net.Listen("tcp", ":"+*thisPort)
	defer listener.Close()
	rpc.Accept(listener)

	if *bottles != 0 {
		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer. Take one down, pass it around...\n", *bottles, *bottles)
		client, _ := rpc.Dial("tcp", nextAddr)
		defer client.Close()

		client.Go()
		client.Call()
	}
}
