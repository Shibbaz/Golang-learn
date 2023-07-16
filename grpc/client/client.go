package main

import (
	pb "apikiller"
	"context"
	"crypto/rand"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func ID(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

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

	Id := ID(5)
	r, err := c.CreateUser(ctx, &pb.UserArgs{
		Id:      Id,
		Name:    "Kamil",
		Surname: "Mosciszko",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf(r.Message)
	r, err = c.GetAllUsers(ctx, &pb.APIRequest{})
	if err != nil {
		log.Fatalf("could not get all users: %v", err)
	}
	log.Printf(r.Message)

	r, err = c.GetUserById(ctx, &pb.UserArgs{Id: Id})
	if err != nil {
		log.Fatalf("could not get user by id: %v", err)
	}
	log.Printf(r.Message)

	r, err = c.UpdateUser(ctx, &pb.UserArgs{Id: Id, Name: "Lukasz", Surname: "Mazurek"})
	if err != nil {
		log.Fatalf("could not update: %v", err)
	}
	log.Printf(r.Message)

	r, err = c.GetUserById(ctx, &pb.UserArgs{Id: Id})
	if err != nil {
		log.Fatalf("could not get user by id: %v", err)
	}
	log.Printf(r.Message)

	r, err = c.DeleteUser(ctx, &pb.UserArgs{Id: Id})
	if err != nil {
		log.Fatalf("could not delete: %v", err)
	}
	log.Printf(r.Message)
}
