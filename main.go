package main

import (
	"io"
	"log"
	"net/http"
	"time"

	pb "github.com/mralanlee/hello-world-protobuf/messages"
)

func myHandler(w http.ResponseWriter, req *http.Request) {
	hello := pb.Hello{
		World: "World",
	}
	io.WriteString(w, hello.World)
}

func main() {
	// Create Server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(myHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
