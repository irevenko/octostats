# octostats 🐙🐱📦
> A supplementary Go package on top of <a href="https://github.com/google/go-github">go-github</a> and <a href="https://github.com/shurcooL/githubv4">githubv4</a>

<p align="center"><img src="octo-gopher.png" width="300"></p>

<p align="center">This package allows you to fetch extra GitHub data</p> <br>

# Methods 🧰
* [REST](#REST "Goto #REST")
	* [AllRepos](#AllRepos "Goto ##AllRepos")
	* [LanguagesByRepo](#LanguagesByRepo "Goto ##LanguagesByRepo")
	* [MostUsedLicenses](#MostUsedLicenses "Goto ##MostUsedLicenses")
	* [MostStarredRepos](#MostStarredRepos "Goto ##MostStarredRepos")
	* [MostForkedRepos](#MostForkedRepos "Goto ##MostForkedRepos")
	* [StarsPerLanguage](#StarsPerLanguage "Goto ##StarsPerLanguage")
	* [ForksPerLanguage](#AllRepos "Goto ##ForksPerLanguage")
	* [TotalStars](#TotalStars "Goto ##TotalStars")
	* [TotalForks](#TotalForks "Goto ##TotalForks")
* [GraphQL](#GraphQL "Goto #GraphQL")
	* [AllContributions](#AllContributions "Goto ##AllContributions")
	* [AllCommits](#AllCommits "Goto ##AllCommits")
	* [AllIssues](#AllIssues "Goto ##AllIssues")
	* [AllPullRequests](#AllPullRequests "Goto ##AllPullRequests")
	* [LanguagesByCommit](#LanguagesByCommit "Goto ##LanguagesByCommit")


# Docs 📋
# REST
All examples are using ```AuthREST``` <br>
```ctx, client := r.AuthREST("<YOUR_TOKEN>")```

If you want you can write your own auth but keep in mind that you in order to use this package ```client, context``` are required


## AllRepos
Returns slice of repos for user/organization (https://api.github.com/users/USERNAME/repos)
```go
import ( 
    "fmt"

    gh "github.com/irevenko/octostats/github"
    "github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README

	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name
	fmt.Println(allRepos)
}
```

## MostUsedLanguages
Returns two slices of names and occurrences
```go
import ( 
	"fmt"
	"strconv"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name

	usedLangs, langsNum := gh.MostUsedLanguages(ctx, client, allRepos)
	fmt.Println("Most used langs")
	for i, v := range usedLangs {
		fmt.Println(v + ": " + strconv.Itoa(langsNum[i]))
	}
}
```

## MostUsedLicenses
Returns two slices of names and occurrences
``` go
import ( 
	"fmt"
	"strconv"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME")// could be user or organization name

	usedLicenses, licsNum := gh.MostUsedLicenses(ctx, client, allRepos)
	fmt.Println("Most used licenses")
	for i, v := range usedLicenses {
		fmt.Println(v + ": " + strconv.Itoa(licsNum[i]))
	}
}
```

## MostStarredRepos
Returns two slices of names and stars num
``` go
import ( 
	"fmt"
	"strconv"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name
    
	starredRepos, starredNums := gh.MostStarredRepos(ctx, client, allRepos)
	fmt.Println("Most starred repos")
	for i, v := range starredRepos {
		fmt.Println(v + ": " + strconv.Itoa(starredNums[i]))
	}
}
```

## MostForkedRepos
Returns two slices of names and forks num
```go
import ( 
	"fmt"
	"strconv"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name
    
	forkedRepos, forkedNums := gh.MostForkedRepos(ctx, client, allRepos)
	fmt.Println("Most forked repos")
	for i, v := range forkedRepos {
		fmt.Println(v + ": " + strconv.Itoa(forkedNums[i]))
	}
}
```

## StarsPerLanguage
Returns two slices of languages and stars num
```go
import ( 
	"fmt"
	"strconv"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name
    
	starsPerL, starsNum := gh.StarsPerLanguage(ctx, client, allRepos)
	fmt.Println("Stars per lang")
	for i, v := range starsPerL {
		fmt.Println(v + ": " + strconv.Itoa(starsNum[i]))
	}
}
```
## ForksPerLanguage
Returns two slices of languages and forks num
```go
import ( 
 	"fmt"
	"strconv"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name

	forksPerL, forksNum := gh.ForksPerLanguage(ctx, client, allRepos)
	fmt.Println("Forks per lang")
	for i, v := range forksPerL {
		fmt.Println(v + ": " + strconv.Itoa(forksNum[i]))
	}
}
```

## TotalStars
Returns integer number
```go
import ( 
	"fmt"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name
    
	totalStars := g.TotalStars(ctx, client, allRepos)
	fmt.Println("Total stars")
	fmt.Println(totalStars)
}
```

## TotalForks
Returns integer number
```go
import ( 
	"fmt"

	gh "github.com/irevenko/octostats/github"
	"github.com/google/go-github/github"
)

func main() {
	ctx, client := auth() //this function is located at the top of the irevenko/octostats Docs section in README
	allRepos := gh.AllRepos(ctx, client, "USERNAME") //could be user or organization name
    
	totalForks := g.TotalForks(ctx, client, allRepos)
	fmt.Println("Total forks")
	fmt.Println(totalForks)
}
```

# GraphQL
All examples are using ```AuthGraphQL``` <br>
```client := r.AuthGraphQL("<YOUR_TOKEN>")```

If you want you can write your own auth but keep in mind that you in order to use this package ```client``` is required


# Contributing 🤝
Contributions, issues and feature requests are welcome! 👍 <br>
Feel free to check [open issues](https://github.com/irevenko/octostats/issues).

# Notes
- shows private repos and repos from orgs when using empty string as name (if authorized)
- see readme-stats, metrics

# License 📑 
(c) 2021 Ilya Revenko. [MIT License](https://tldrlegal.com/license/mit-license)
