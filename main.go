package main

import (
	"context"
	"fmt"

	g "./github"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	allRepos := g.GetAllRepos(ctx, client, "skanehira")

	languages, numbers := g.GetMostForkedRepos(ctx, client, allRepos)
	fmt.Println("Most forked repos")
	fmt.Println(languages)
	fmt.Println(numbers)

	languages1, numbers1 := g.GetMostStarredRepos(ctx, client, allRepos)
	fmt.Println("Most starred repos")
	fmt.Println(languages1)
	fmt.Println(numbers1)

	languages2, numbers2 := g.GetMostUsedLanguages(ctx, client, allRepos)
	fmt.Println("Most used langs")
	fmt.Println(languages2)
	fmt.Println(numbers2)

	langsS, stars := g.GetStarsPerLanguage(ctx, client, allRepos)
	fmt.Println("Stars per lang")
	fmt.Println(langsS)
	fmt.Println(stars)

	langsF, forks := g.GetForksPerLanguage(ctx, client, allRepos)
	fmt.Println("Forks per lang")
	fmt.Println(langsF)
	fmt.Println(forks)

	totalStars := g.GetTotalStars(ctx, client, allRepos)
	fmt.Println("Total stars")
	fmt.Println(totalStars)

	totalForks := g.GetTotalForks(ctx, client, allRepos)
	fmt.Println("Total forks")
	fmt.Println(totalForks)

	fmt.Println("Total repos")
	fmt.Println(len(allRepos))
}
