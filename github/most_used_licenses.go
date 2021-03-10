package github

import (
	"context"

	h "../helpers"
	"github.com/google/go-github/github"
)

func MostUsedLicenses(ctx context.Context, client *github.Client, allRepos []*github.Repository) (licenses []string, occurrences []int) {
	var licensesSlice []string

	for _, v := range allRepos {
		if v.License != nil {
			license := *v.License.Name
			licensesSlice = append(licensesSlice, license)
		} else {
			license := "No License"
			licensesSlice = append(licensesSlice, license)
		}
	}

	mostUsedLicenses := h.CountDuplicates(licensesSlice)
	lics, occurrs := h.SortMap(mostUsedLicenses)

	return lics, occurrs
}
