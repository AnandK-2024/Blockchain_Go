package main

import (
	"fmt"
	"time"
	// "github.com/AnandK-2024/Blockchain/network"
)

// func main() {
// 	fmt.Println("Hello, World!")
// 	//  Create a new LocalTransport instance
// 	transport := network.NewLocaltransport(network.NetAddr("localhost"))
// 	opts := network.serveropts{
// 		transport: []network.Transport{transport} ,
// 	}
// 	network.Newserver(opts)

// }

func printMessage(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}

func main() {
	defer go printMessage("Hello")
	go printMessage("World")

	// Wait for goroutine to complete
	time.Sleep(5 * time.Second)
}
