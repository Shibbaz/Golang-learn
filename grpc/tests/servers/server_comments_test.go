package servers

import (
	"context"
	"testing"

	. "comments"
	. "helpers"
	. "posts"
	. "users"
)

func TestCreateComment(t *testing.T) {
	ctx := context.Background()
	userClient, userCloser := NewUsersTestServer(ctx)
	postClient, postCloser := NewPostsTestServer(ctx)
	client, closer := NewCommentsTestServer(ctx)
	defer userCloser()
	defer postCloser()
	defer closer()

	id := Generate(6)
	userId := Generate(6)
	postId := Generate(6)

	type expectation struct {
		out *CommentsReply
		err error
	}
	tests := map[string]struct {
		in       *CommentArgs
		expected expectation
	}{
		"Must_Success": {
			in: &CommentArgs{
				Id:       id,
				Text:     "kamil44",
				AuthorId: userId,
				PostId:   postId,
			},
			expected: expectation{
				out: &CommentsReply{
					Message: "[{id:" + id + " text:kamil44, published:false}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := userClient.CreateUser(ctx, &UserArgs{
				Id:      userId,
				Name:    "kamil44",
				Surname: "Mos44",
			})
			if err != nil {
				panic(err)
			}
			_, err = postClient.CreatePost(ctx, &PostArgs{
				Id:       postId,
				Title:    "Kamil",
				AuthorId: userId,
			})
			if err != nil {
				panic(err)
			}
			comment, err := client.CreateComment(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != "[{id:"+id+" text:kamil44}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, comment)
			}
		})
	}
}
