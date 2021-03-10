package main

import (
	"context"
	"fmt"
	"strconv"

	g "./github"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func auth() (context.Context, *github.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	return ctx, client
}

func main() {
	ctx, client := auth()

	allRepos := g.AllRepos(ctx, client, "hajimehoshi")

	forkedRepos, forkedNums := g.MostForkedRepos(ctx, client, allRepos)
	fmt.Println("Most forked repos")
	for i, v := range forkedRepos {
		fmt.Println(v + ": " + strconv.Itoa(forkedNums[i]))
	}

	starredRepos, starredNums := g.MostStarredRepos(ctx, client, allRepos)
	fmt.Println("\nMost starred repos")
	for i, v := range starredRepos {
		fmt.Println(v + ": " + strconv.Itoa(starredNums[i]))
	}

	usedLangs, langsNum := g.MostUsedLanguages(ctx, client, allRepos)
	fmt.Println("\nMost used langs")
	for i, v := range usedLangs {
		fmt.Println(v + ": " + strconv.Itoa(langsNum[i]))
	}

	usedLicenses, licsNum := g.MostUsedLicenses(ctx, client, allRepos)
	fmt.Println("\nMost used licenses")
	for i, v := range usedLicenses {
		fmt.Println(v + ": " + strconv.Itoa(licsNum[i]))
	}

	starsPerL, starsNum := g.StarsPerLanguage(ctx, client, allRepos)
	fmt.Println("\nStars per lang")
	for i, v := range starsPerL {
		fmt.Println(v + ": " + strconv.Itoa(starsNum[i]))
	}

	forksPerL, forksNum := g.ForksPerLanguage(ctx, client, allRepos)
	fmt.Println("\nForks per lang")
	for i, v := range forksPerL {
		fmt.Println(v + ": " + strconv.Itoa(forksNum[i]))
	}

	totalStars := g.TotalStars(ctx, client, allRepos)
	fmt.Println("\nTotal stars")
	fmt.Println(totalStars)

	totalForks := g.TotalForks(ctx, client, allRepos)
	fmt.Println("\nTotal forks")
	fmt.Println(totalForks)

	fmt.Println("\nTotal repos")
	fmt.Println(len(allRepos))
}
