package server

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	prisma "prisma_client"

	. "comments"
	. "posts"
	. "users"
)

var (
	Port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) GetAllUsers(ctx context.Context, in *APIRequest) (*UserReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()
	users, err := Client.Users(nil).Exec(Context)
	if err != nil {
		panic(err)
	}
	if in.Page > 0 && in.Limit > 0 {
		skip := in.Page * in.Limit
		users, err = Client.Users(&prisma.UsersParams{
			First: &in.Limit,
			Skip:  &skip,
		}).Exec(Context)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("GetAllUsers resolver %+v\n", users)

	out, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	return &UserReply{Message: string(out)}, nil
}

func (s *server) GetUserById(ctx context.Context, in *UserArgs) (*UserReply, error) {
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
	return &UserReply{Message: string(out)}, nil
}

func (s *server) CreateUser(ctx context.Context, in *UserArgs) (*UserReply, error) {
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
	return &UserReply{Message: string(out)}, nil
}

func (s *server) DeleteUser(ctx context.Context, in *UserArgs) (*UserReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()

	_, err := Client.DeleteUser(prisma.UserWhereUniqueInput{
		ID: &in.Id,
	}).Exec(Context)

	if err != nil {
		panic(err)
	}
	message := fmt.Sprintf("Deleted user %s", in.Id)
	fmt.Printf(message)
	return &UserReply{Message: message}, nil
}

func (s *server) UpdateUser(ctx context.Context, in *UserArgs) (*UserReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()
	user, err := Client.UpdateUser(prisma.UserUpdateParams{
		Where: prisma.UserWhereUniqueInput{
			ID: &in.Id,
		},
		Data: prisma.UserUpdateInput{
			Name:    &in.Name,
			Surname: &in.Surname,
		},
	}).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated user: %+v\n", user)
	out, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return &UserReply{Message: string(out)}, nil
}

func (s *server) CreatePost(ctx context.Context, in *PostArgs) (*PostsReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()
	post, err := Client.CreatePost(prisma.PostCreateInput{
		ID:          &in.Id,
		Title:       in.Title,
		Description: &in.Description,
		Author: &prisma.UserCreateOneInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &in.AuthorId,
			},
		},
	},
	).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create post: %+v\n", post)

	out, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	return &PostsReply{Message: string(out)}, nil
}

func (s *server) DeletePost(ctx context.Context, in *PostArgs) (*PostsReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()

	_, err := Client.DeletePost(prisma.PostWhereUniqueInput{
		ID: &in.Id,
	}).Exec(Context)

	if err != nil {
		panic(err)
	}
	message := fmt.Sprintf("Deleted post %s", in.Id)
	fmt.Printf(message)
	return &PostsReply{Message: message}, nil
}

func (s *server) UpdatePost(ctx context.Context, in *PostArgs) (*PostsReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()
	post, err := Client.UpdatePost(prisma.PostUpdateParams{
		Where: prisma.PostWhereUniqueInput{
			ID: &in.Id,
		},
		Data: prisma.PostUpdateInput{
			Title:       &in.Title,
			Description: &in.Description,
		},
	}).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated post: %+v\n", post)
	out, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}
	return &PostsReply{Message: string(out)}, nil
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateComment(ctx context.Context, in *CommentArgs) (*CommentsReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()
	comment, err := Client.CreateComment(prisma.CommentCreateInput{
		ID:   &in.Id,
		Text: &in.Text,
		Author: &prisma.UserCreateOneInput{
			Connect: &prisma.UserWhereUniqueInput{
				ID: &in.AuthorId,
			},
		},
		Post: &prisma.PostCreateOneInput{
			Connect: &prisma.PostWhereUniqueInput{
				ID: &in.PostId,
			},
		},
	},
	).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Create comment: %+v\n", comment)

	out, err := json.Marshal(comment)
	if err != nil {
		panic(err)
	}
	return &CommentsReply{Message: string(out)}, nil
}

func (s *server) DeleteComment(ctx context.Context, in *CommentArgs) (*CommentsReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()

	_, err := Client.DeleteComment(prisma.CommentWhereUniqueInput{
		ID: &in.Id,
	}).Exec(Context)

	if err != nil {
		panic(err)
	}
	message := fmt.Sprintf("Deleted comment %s", in.Id)
	fmt.Printf(message)
	return &CommentsReply{Message: message}, nil
}

func (s *server) UpdateComment(ctx context.Context, in *CommentArgs) (*CommentsReply, error) {
	Client := prisma.New(nil)
	Context := context.TODO()
	comment, err := Client.UpdateComment(prisma.CommentUpdateParams{
		Where: prisma.CommentWhereUniqueInput{
			ID: &in.Id,
		},
		Data: prisma.CommentUpdateInput{
			Text: &in.Text,
		},
	}).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated comment: %+v\n", comment)
	out, err := json.Marshal(comment)
	if err != nil {
		panic(err)
	}
	return &CommentsReply{Message: string(out)}, nil
}
