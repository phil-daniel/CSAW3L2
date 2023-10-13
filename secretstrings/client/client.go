package main

import (
	"flag"
	"net/rpc"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"bufio"
	//	"os"
	"fmt"
)

func main() {
	server := flag.String("server", "127.0.0.1:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	//TODO: connect to the RPC server and send the request(s)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	request := stubs.Request{Message: "Hello"}
	response := new(stubs.Response)

	client.Call(stubs.ReverseHandler, request, response)
	fmt.Println("Responded: " + response.Message)

	request1 := stubs.Request{Message: "Hello"}
	response1 := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request1, response1)
	fmt.Println("Fast response: " + response1.Message)
}
