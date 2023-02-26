package main

import (
	"context"

	"github.com/tMinamiii/go-http-server/router"
	"github.com/tMinamiii/go-http-server/server"
)

func main() {
	ctx := context.Background()
	r := router.Router()
	s := server.NewServer(r)
	s.Run(ctx)
}
