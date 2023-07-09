package main

import (
	"context"

	"github.com/shubhashreeb/ddos/pkg"
)

func main() {
	ctx := context.Background()
	go pkg.RunGrpcServer(ctx)
	go func() {
		_ = pkg.RunRestServer(ctx)
	}()
	//return pkg.RunGrpcServer(ctx)

	select {}
}
