package rest

import (
	"github.com/google/go-github/v33/github"
)

func TotalForks(client *github.Client, allRepos []*github.Repository) (forksNum int) {
	_, forks := ForksPerLanguage(client, allRepos)

	var totalForks int
	for _, v := range forks {
		totalForks += v
	}

	return totalForks
}
