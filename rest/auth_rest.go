package rest

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func AuthREST(token string) (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return ctx, client
}
