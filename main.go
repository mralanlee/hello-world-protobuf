package main

import (
	"fmt"

	pb "github.com/mralanlee/hello-world-protobuf/messages"
)

func main() {
	fmt.Println("Hello")
	hello := pb.Hello{
		World: "World",
	}
	fmt.Printf("Hello %s", hello.World)
}
