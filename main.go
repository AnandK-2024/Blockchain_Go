package main

import (
	"fmt"

	"github.com/AnandK-2024/Blockchain/network"
	// "github.com/AnandK-2024/Blockchain/network"
)

func main() {
	fmt.Println("Hello, World!")
	//  Create a new LocalTransport instance
	transport := network.NewLocaltransport(network.NetAddr("localhost"))
	opts := network.serveropts{
		transport: []network.Transport{transport},
	}
	s := network.Newserver(opts)
	s.start()

}
