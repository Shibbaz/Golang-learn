package prisma

import (
	"context"
	"fmt"

	prisma "prisma-client"
)

func main() {
	Client := prisma.New(nil)
	Context := context.TODO()

	// Create a new user
	name := "Alice"
	newUser, err := Client.CreateUser(prisma.UserCreateInput{
		Name: name,
	}).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created new user: %+v\n", newUser)

	users, err := Client.Users(nil).Exec(Context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", users)
}
