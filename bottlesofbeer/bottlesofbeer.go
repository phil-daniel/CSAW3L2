package main

import (
	"flag"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"time"
	//	"net/rpc"
	//	"fmt"
	//	"time"
	//	"net"
)

var nextAddr string

type Shared int

// Calls the nextAddr, telling it to close the program. If the next address is no longer up it automatically closes
func (t *Shared) Exit(status *int, stopAt *string) error {
	client, err := rpc.Dial("tcp", nextAddr)
	if err != nil {
		os.Exit(*status)
	}
	defer client.Close()
	client.Go("Shared.Exit", 0, 0, nil)
	os.Exit(*status)
	return nil
}

func (t *Shared) Bottles(bottles *int, reply *int) error {
	if *bottles > 0 {
		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer. Take one down, pass it around...\n\n", *bottles, *bottles)
		time.Sleep(50 * time.Millisecond)
		client, err := rpc.Dial("tcp", nextAddr)
		if err != nil {
			fmt.Println("Something went wrong with the nextAddr")
			os.Exit(1)
		}
		defer client.Close()
		client.Go("Shared.Bottles", *bottles-1, 0, nil)
	} else if *bottles == 0 {
		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer. No more bottles of beer...\n\n", *bottles, *bottles)

		client, err := rpc.Dial("tcp", nextAddr)
		if err != nil {
			fmt.Println("Something went wrong with the nextAddr")
			os.Exit(1)
		}
		defer client.Close()
		client.Go("Shared.Exit", 0, 0, nil)
	}
	return nil
}

func main() {
	thisPort := flag.String("this", "8030", "Port for this process to listen on")
	flag.StringVar(&nextAddr, "next", "localhost:8040", "IP:Port string for next member of the round.")
	bottles := flag.Int("n", 0, "Bottles of Beer (launches song if not 0)")
	flag.Parse()
	//TODO: Up to you from here! Remember, you'll need to both listen for
	//RPC calls and make your own.

	shared := new(Shared)
	rpc.Register(shared)

	// Initial call if bottles wasn't 0 in the input, this will launch the song
	if *bottles != 0 {
		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer. Take one down, pass it around...\n\n", *bottles, *bottles)
		// This is next address to call
		client, err := rpc.Dial("tcp", nextAddr)
		if err != nil {
			fmt.Println("Something went wrong with the nextAddr")
			os.Exit(1)
		}
		defer client.Close()
		client.Go("Shared.Bottles", *bottles-1, 0, nil)
	}

	// Setting up listening
	listener, _ := net.Listen("tcp", ":"+*thisPort)
	defer listener.Close()
	rpc.Accept(listener)
}
