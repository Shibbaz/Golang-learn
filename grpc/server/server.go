package server

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	pb "grpc"
	"prisma-client"
)

var (
	Port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) GetAllUsers(ctx context.Context, in *pb.APIRequest) (*pb.APIReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()

	users, err := Client.Users(nil).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetAllUsers resolver %+v\n", users)

	out, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	return &pb.APIReply{Message: string(out)}, nil
}

func (s *server) GetUserById(ctx context.Context, in *pb.UserArgs) (*pb.APIReply, error) {
	Client := prisma.New(nil)
	user, err := Client.Users(&prisma.UsersParams{
		Where: &prisma.UserWhereInput{
			ID: &in.Id,
		},
	}).Exec(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Printf("GetUserById resolver %+v\n", user)

	out, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return &pb.APIReply{Message: string(out)}, nil
}

func (s *server) CreateUser(ctx context.Context, in *pb.UserArgs) (*pb.APIReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()

	newUser, err := Client.CreateUser(prisma.UserCreateInput{
		ID:      &in.Id,
		Name:    in.Name,
		Surname: in.Surname,
	}).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created new user: %+v\n", newUser)
	out, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}
	return &pb.APIReply{Message: string(out)}, nil
}
func NewServer() *server {
	return &server{}
}
