package main

import (
	"context"
	"log"

	"github.com/mec07/demo-grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// create client TLS creds
	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load TLS certs: %s", err)
	}

	conn, err := grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewPingClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)
}
