package server

import (
	"context"
	"crypto/rand"
	"fmt"
	pb "grpc"
	"log"
	"net"
	"testing"
	"unsafe"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

var alphabet = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generate(size int) string {
	b := make([]byte, size)
	rand.Read(b)
	for i := 0; i < size; i++ {
		b[i] = alphabet[b[i]%byte(len(alphabet))]
	}
	return *(*string)(unsafe.Pointer(&b))
}

func NewTestServer(ctx context.Context) (pb.APIClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	pb.RegisterAPIServer(baseServer, NewServer())
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := pb.NewAPIClient(conn)

	return client, closer
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	client, closer := NewTestServer(ctx)
	defer closer()
	id := generate(6)
	type expectation struct {
		out *pb.APIReply
		err error
	}

	tests := map[string]struct {
		in       *pb.UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &pb.UserArgs{
				Id:      id,
				Name:    "kamil4",
				Surname: "Mos4",
			},
			expected: expectation{
				out: &pb.APIReply{
					Message: "[{ID:" + id + " Name:kamil Surname:Mos}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			users, err := client.CreateUser(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != "[{ID:"+id+" Name:kamil Surname:Mos}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, users)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	ctx := context.Background()

	client, closer := NewTestServer(ctx)
	defer closer()

	type expectation struct {
		out *pb.APIReply
		err error
	}
	tests := map[string]struct {
		in       *pb.UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &pb.UserArgs{
				Name:    "kamil4",
				Surname: "Mos4",
			},
			expected: expectation{
				out: &pb.APIReply{
					Message: "[{Name:kamil4 Surname:Mos4}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			users, err := client.GetUserById(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message == "[{Name:kamil Surname:Mos}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, users)
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()

	client, closer := NewTestServer(ctx)
	defer closer()
	id := generate(6)
	type expectation struct {
		out *pb.APIReply
		err error
	}
	message := fmt.Sprintf("Deleted user %s", id)

	tests := map[string]struct {
		in       *pb.UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &pb.UserArgs{
				Id: id,
			},
			expected: expectation{
				out: &pb.APIReply{
					Message: message,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.CreateUser(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			_, err = client.DeleteUser(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != message {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, message)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()

	client, closer := NewTestServer(ctx)
	defer closer()
	id := generate(6)
	type expectation struct {
		out *pb.APIReply
		err error
	}

	tests := map[string]struct {
		in       *pb.UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &pb.UserArgs{
				Id:      id,
				Name:    "kamil44",
				Surname: "Mos44",
			},
			expected: expectation{
				out: &pb.APIReply{
					Message: "[{ID:" + id + " Name:kamil44 Surname:Mos44}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.CreateUser(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			user, err := client.UpdateUser(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != "[{ID:"+id+" Name:kamil44 Surname:Mos44}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, user)
			}
		})
	}
}
