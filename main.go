package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tMinamiii/go-http-server/config"
	"github.com/tMinamiii/go-http-server/server"
)

func newServer(ctx context.Context) *server.Server {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to open config %v", err)
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}

	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	mux := server.NewMux()
	s := server.NewServer(l, mux)
	return s
}

func main() {
	ctx := context.Background()
	s := newServer(ctx)
	go s.Run(ctx)
	s.GracefulShutdown()
}
