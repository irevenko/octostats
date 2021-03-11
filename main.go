package main

import (
	"fmt"

	g "./graphql"
	r "./rest"
)

func main() {
	ctx, client := r.AuthREST("")
	qlClient := g.AuthGraphQL("")

	allRepos := r.AllRepos(ctx, client, "skanehira")

	l, n := r.MostStarredRepos(client, allRepos)
	fmt.Println(l[:10])
	fmt.Println(n[:10])

	fmt.Println(g.AllContributions(qlClient, "skanehira", 2020, 2021))

	//g.LanguagesByCommit(qlClient, "denbondd", 2020, 2021)

	//langs := g.LanguagesByCommit(qlClient, "irevenko", 2020, 2021)

	//fmt.Println(commits)

	// forkedRepos, forkedNums := r.MostForkedRepos(client, allRepos)
	// fmt.Println("Most forked repos")
	// for i, v := range forkedRepos {
	// 	fmt.Println(v + ": " + strconv.Itoa(forkedNums[i]))
	// }

	// starredRepos, starredNums := r.MostStarredRepos(client, allRepos)
	// fmt.Println("\nMost starred repos")
	// for i, v := range starredRepos {
	// 	fmt.Println(v + ": " + strconv.Itoa(starredNums[i]))
	// }

	// usedLangs, langsNum := r.LangsByRepo(client, allRepos)
	// fmt.Println("\nMost used langs")
	// for i, v := range usedLangs {
	// 	fmt.Println(v + ": " + strconv.Itoa(langsNum[i]))
	// }

	// usedLicenses, licsNum := r.MostUsedLicenses(client, allRepos)
	// fmt.Println("\nMost used licenses")
	// for i, v := range usedLicenses {
	// 	fmt.Println(v + ": " + strconv.Itoa(licsNum[i]))
	// }

	// starsPerL, starsNum := r.StarsPerLanguage(client, allRepos)
	// fmt.Println("\nStars per lang")
	// for i, v := range starsPerL {
	// 	fmt.Println(v + ": " + strconv.Itoa(starsNum[i]))
	// }

	// forksPerL, forksNum := r.ForksPerLanguage(client, allRepos)
	// fmt.Println("\nForks per lang")
	// for i, v := range forksPerL {
	// 	fmt.Println(v + ": " + strconv.Itoa(forksNum[i]))
	// }

	// totalStars := r.TotalStars(client, allRepos)
	// fmt.Println("\nTotal stars")
	// fmt.Println(totalStars)

	// totalForks := r.TotalForks(client, allRepos)
	// fmt.Println("\nTotal forks")
	// fmt.Println(totalForks)

	// fmt.Println("\nTotal repos")
	// fmt.Println(len(allRepos))
}
