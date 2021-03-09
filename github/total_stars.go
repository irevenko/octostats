package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetTotalStars(ctx context.Context, client *github.Client, allRepos []*github.Repository) int {
	_, stars := GetStarsPerLanguage(ctx, client, allRepos)

	var totalStars int
	for _, v := range stars {
		totalStars += v
	}

	return totalStars
}
