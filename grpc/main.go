package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"posts"
	"server"
	"users"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a listener on TCP port
	usersListen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	postsListen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	serv := server.NewServer()
	// Create a gRPC server object
	usersServer := grpc.NewServer()
	postsServer := grpc.NewServer()

	// Attach the Greeter service to the server
	users.RegisterUsersAPIServer(usersServer, serv)
	posts.RegisterPostsAPIServer(postsServer, serv)

	// Serve gRPC server
	log.Println("Serving users/gRPC on 0.0.0.0:8080")
	log.Println("Serving posts/gRPC on 0.0.0.0:8081")

	go func() {
		log.Fatalln(usersServer.Serve(usersListen))

	}()
	go func() {
		log.Fatalln(postsServer.Serve(postsListen))
	}()

	// Create a client connection to the gRPC server we just started
	usersConn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	postsConn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8081",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	// Create a new ServeMux for the gRPC-Gateway
	gwmux := runtime.NewServeMux()
	// Register the Greeter service with the gRPC-Gateway
	err = users.RegisterUsersAPIHandler(context.Background(), gwmux, usersConn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	err = posts.RegisterPostsAPIHandler(context.Background(), gwmux, postsConn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Create a new HTTP server for the gRPC-Gateway
	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
