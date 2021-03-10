package github

import (
	"context"

	"github.com/google/go-github/github"
)

func TotalForks(ctx context.Context, client *github.Client, allRepos []*github.Repository) (forksNum int) {
	_, forks := ForksPerLanguage(ctx, client, allRepos)

	var totalForks int
	for _, v := range forks {
		totalForks += v
	}

	return totalForks
}
