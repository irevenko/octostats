package rest

import (
	"github.com/google/go-github/v33/github"
)

func TotalStars(client *github.Client, allRepos []*github.Repository) (starsNum int) {
	_, stars := StarsPerLanguage(client, allRepos)

	var totalStars int
	for _, v := range stars {
		totalStars += v
	}

	return totalStars
}
