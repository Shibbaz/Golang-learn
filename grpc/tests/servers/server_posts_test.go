package servers

import (
	"context"
	prisma "prisma_client"
	"testing"

	. "helpers"
	. "posts"
	. "users"
)

func TestCreatePost(t *testing.T) {
	ctx := context.Background()
	userClient, userCloser := NewUsersTestServer(ctx)
	client, closer := NewPostsTestServer(ctx)
	defer closer()
	defer userCloser()

	id := Generate(6)
	type expectation struct {
		out *PostsReply
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
				out: &PostsReply{
					Message: "[{ID:" + id + " Title:Kamil}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := userClient.CreateUser(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			post, err := client.CreatePost(ctx, &PostArgs{
				Title:    "Kamil",
				AuthorId: id,
			})
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != "[{ID:"+id+" Title:Kamil}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, post)
			}
		})
	}
}

func TestUpdatePost(t *testing.T) {
	ctx := context.Background()

	client, closer := NewPostsTestServer(ctx)
	defer closer()

	type expectation struct {
		out *PostsReply
		err error
	}
	description := "X"
	post, _ := prisma.New(nil).CreatePost(prisma.PostCreateInput{
		Title:       "Kamil",
		Description: &description,
	}).Exec(ctx)
	PostId := post.ID
	tests := map[string]struct {
		in       *PostArgs
		expected expectation
	}{
		"Must_Success": {
			in: &PostArgs{
				Id:          PostId,
				Title:       "Kamils",
				Description: "XD",
			},
			expected: expectation{
				out: &PostsReply{
					Message: "[{ID:" + PostId + " Title:Kamils Description:XD}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			post, err := client.UpdatePost(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != "[{ID:"+PostId+" Title:Kamils Description:XD}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, post)
			}
		})
	}
}
