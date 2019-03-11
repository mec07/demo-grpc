package api

import (
	context "context"
	"log"
)

// Server represents the gRPC server
type Server struct{}

// SayHello generates a response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *PingMessage) (*PingMessage, error) {
	log.Printf("Received message %s", in.Greeting)
	return &PingMessage{Greeting: "bar"}, nil
}
