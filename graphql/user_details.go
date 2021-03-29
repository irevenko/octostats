package graphql

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func UserDetails(client *githubv4.Client, user string) (innerUser User, err error) {
	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	clientErr := client.Query(context.Background(), &UserQuery, variables)
	if clientErr != nil {
		return innerUser, fmt.Errorf("Couldn't get user %s: %w", user, clientErr)
	}

	return UserQuery.User, nil
}
