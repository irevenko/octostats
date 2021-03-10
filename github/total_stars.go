package github

import (
	"context"

	"github.com/google/go-github/github"
)

func TotalStars(ctx context.Context, client *github.Client, allRepos []*github.Repository) (starsNum int) {
	_, stars := StarsPerLanguage(ctx, client, allRepos)

	var totalStars int
	for _, v := range stars {
		totalStars += v
	}

	return totalStars
}
