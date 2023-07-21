package servers

import (
	"context"
	"fmt"
	"testing"

	. "helpers"
	. "users"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	client, closer := NewUsersTestServer(ctx)
	defer closer()
	id := Generate(6)
	type expectation struct {
		out *UserReply
		err error
	}

	tests := map[string]struct {
		in       *UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &UserArgs{
				Id:      id,
				Name:    "kamil4",
				Surname: "Mos4",
			},
			expected: expectation{
				out: &UserReply{
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

	client, closer := NewUsersTestServer(ctx)
	defer closer()

	type expectation struct {
		out *UserReply
		err error
	}
	tests := map[string]struct {
		in       *UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &UserArgs{
				Name:    "kamil4",
				Surname: "Mos4",
			},
			expected: expectation{
				out: &UserReply{
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

	client, closer := NewUsersTestServer(ctx)
	defer closer()
	id := Generate(6)
	type expectation struct {
		out *UserReply
		err error
	}
	message := fmt.Sprintf("Deleted user %s", id)

	tests := map[string]struct {
		in       *UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &UserArgs{
				Id: id,
			},
			expected: expectation{
				out: &UserReply{
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

	client, closer := NewUsersTestServer(ctx)
	defer closer()
	id := Generate(6)
	type expectation struct {
		out *UserReply
		err error
	}

	tests := map[string]struct {
		in       *UserArgs
		expected expectation
	}{
		"Must_Success": {
			in: &UserArgs{
				Id:      id,
				Name:    "kamil44",
				Surname: "Mos44",
			},
			expected: expectation{
				out: &UserReply{
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
