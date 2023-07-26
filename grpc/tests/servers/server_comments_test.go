package servers

import (
	"context"
	"fmt"
	prisma "prisma_client"
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

func TestDeleteComment(t *testing.T) {
	ctx := context.Background()
	client, closer := NewCommentsTestServer(ctx)
	defer closer()

	comment, _ := prisma.New(nil).CreateComment(prisma.CommentCreateInput{}).Exec(ctx)

	type expectation struct {
		out *CommentsReply
		err error
	}
	message := fmt.Sprintf("Deleted comment %s", comment.ID)

	tests := map[string]struct {
		in       *CommentArgs
		expected expectation
	}{
		"Must_Success": {
			in: &CommentArgs{
				Id: comment.ID,
			},
			expected: expectation{
				out: &CommentsReply{
					Message: message,
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			_, err := client.DeleteComment(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != message {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, message)
			}
		})
	}
}

func TestUpdateComment(t *testing.T) {
	ctx := context.Background()

	client, closer := NewCommentsTestServer(ctx)
	defer closer()

	type expectation struct {
		out *CommentsReply
		err error
	}
	comment, _ := prisma.New(nil).CreateComment(prisma.CommentCreateInput{}).Exec(ctx)
	CommentId := comment.ID
	tests := map[string]struct {
		in       *CommentArgs
		expected expectation
	}{
		"Must_Success": {
			in: &CommentArgs{
				Id:   CommentId,
				Text: "Kamilsd",
			},
			expected: expectation{
				out: &CommentsReply{
					Message: "[{ID:" + CommentId + " Title:Kamilsd}]",
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			comment, err := client.UpdateComment(ctx, tt.in)
			if err != nil {
				panic(err)
			}
			if tt.expected.out.Message != "[{ID:"+CommentId+" Text:Kamilsd}]" {
				t.Errorf("Out -> \nWant: %q\nGot: %q", tt.expected.out, comment)
			}
		})
	}
}
