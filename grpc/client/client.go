package main

import (
	"context"
	"flag"
	pb "grpc"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAPIClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CreateUser(ctx, &pb.UserArgs{
		Id:      "XDDDDDDDDD",
		Name:    "Kamil",
		Surname: "Mosciszko",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	r, err = c.GetAllUsers(ctx, &pb.APIRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r.Message)

	r, err = c.GetUserById(ctx, &pb.UserArgs{Id: "1"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r.Message)
}
