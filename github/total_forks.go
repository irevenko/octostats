package github

import (
	"context"

	"github.com/google/go-github/github"
)

func GetTotalForks(ctx context.Context, client *github.Client, allRepos []*github.Repository) int {
	_, forks := GetForksPerLanguage(ctx, client, allRepos)

	var totalForks int
	for _, v := range forks {
		totalForks += v
	}

	return totalForks
}
