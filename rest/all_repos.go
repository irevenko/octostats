package rest

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"
)

func AllRepos(ctx context.Context, client *github.Client, account string) (allRepos []*github.Repository, err error) {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}

	for {
		repos, resp, clientErr := client.Repositories.List(ctx, account, opt)
		if clientErr != nil {
			err = fmt.Errorf("Couldn't get all repositories for %s: %w", account, clientErr)
			break
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return
}
