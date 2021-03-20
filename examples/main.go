package main

import (
	"fmt"
	"strconv"

	"github.com/google/go-github/v33/github"
	g "github.com/irevenko/octostats/graphql"
	r "github.com/irevenko/octostats/rest"
	"github.com/shurcooL/githubv4"
)

func main() {
	// REST auth
	ctx, client := r.AuthREST("<YOUR_TOKEN>")
	// get all repos to work with
	allRepos := r.AllRepos(ctx, client, "<USER_OR_ORGANIZATION>")
	// execute REST examples
	RestExamples(client, allRepos)

	// GraphQL auth
	qlClient := g.AuthGraphQL("<YOUR_TOKEN>")
	// execute GraphQL examples
	GraphqlExamples(qlClient, "<USER_OR_ORGANIZATION>")
}

func RestExamples(client *github.Client, allRepos []*github.Repository) {
	forkedRepos, forkedNums := r.MostForkedRepos(client, allRepos)
	fmt.Println("Most forked repos:")
	for i, v := range forkedRepos {
		fmt.Println(v + ": " + strconv.FormatFloat(forkedNums[i], 'f', -1, 64))
	}

	starredRepos, starredNums := r.MostStarredRepos(client, allRepos)
	fmt.Println("\nMost starred repos:")
	for i, v := range starredRepos {
		fmt.Println(v + ": " + strconv.FormatFloat(starredNums[i], 'f', -1, 64))
	}

	usedLangs, langsNum := r.LanguagesByRepo(client, allRepos)
	fmt.Println("\nLanguages by repo:")
	for i, v := range usedLangs {
		fmt.Println(v + ": " + strconv.FormatFloat(langsNum[i], 'f', -1, 64))
	}

	usedLicenses, licsNum := r.MostUsedLicenses(client, allRepos)
	fmt.Println("\nMost used licenses:")
	for i, v := range usedLicenses {
		fmt.Println(v + ": " + strconv.FormatFloat(licsNum[i], 'f', -1, 64))
	}

	starsPerL, starsNum := r.StarsPerLanguage(client, allRepos)
	fmt.Println("\nStars per lang:")
	for i, v := range starsPerL {
		fmt.Println(v + ": " + strconv.FormatFloat(starsNum[i], 'f', -1, 64))
	}

	forksPerL, forksNum := r.ForksPerLanguage(client, allRepos)
	fmt.Println("\nForks per lang:")
	for i, v := range forksPerL {
		fmt.Println(v + ": " + strconv.FormatFloat(forksNum[i], 'f', -1, 64))
	}

	totalStars := r.TotalStars(client, allRepos)
	fmt.Println("\nTotal stars")
	fmt.Println(totalStars)

	totalForks := r.TotalForks(client, allRepos)
	fmt.Println("\nTotal forks")
	fmt.Println(totalForks)

	fmt.Println("\nTotal repos")
	fmt.Println(len(allRepos))
}

func GraphqlExamples(qlClient *githubv4.Client, user string) {
	langs, commits := g.LanguagesByCommit(qlClient, user, 2020, 2021)
	fmt.Println("\nLanguages by commit 2020-2021:")
	for i, v := range langs {
		fmt.Printf("%v : %v\n", v, commits[i])
	}

	allCommits := g.AllCommits(qlClient, user, 2020, 2021)
	fmt.Println("\nAll commits 2020-2021:")
	fmt.Println(allCommits)

	allPrs := g.AllPullRequests(qlClient, user, 2020, 2021)
	fmt.Println("\nAll pull requests 2020-2021:")
	fmt.Println(allPrs)

	allIssues := g.AllIssues(qlClient, user, 2020, 2021)
	fmt.Println("\nAll issues 2020-2021:")
	fmt.Println(allIssues)

	allContribs := g.AllContributions(qlClient, user, 2020, 2021)
	fmt.Println("\nAll contributions 2020-2021:")
	fmt.Println(allContribs)

	fmt.Println("\nLast year activity:")
	dates, contribs := g.YearActivity(qlClient, user)
	fmt.Println(dates, contribs)

	fmt.Println("\nUser Details:")
	userInfo := g.UserDetails(qlClient, user)
	fmt.Println(userInfo)

	fmt.Println("\nOrganization Details:")
	orgInfo := g.OrganizationDetails(qlClient, "microsoft")
	fmt.Println(orgInfo)
}
