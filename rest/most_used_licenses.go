package rest

import (
	"github.com/google/go-github/v33/github"
	h "github.com/irevenko/octostats/helpers"
)

func MostUsedLicenses(client *github.Client, allRepos []*github.Repository) (licenses []string, occurrences []float64) {
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
