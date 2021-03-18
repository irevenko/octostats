package graphql

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
)

func UserDetails(client *githubv4.Client, user string) User {
	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	err := client.Query(context.Background(), &UserQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return UserQuery.User
}
