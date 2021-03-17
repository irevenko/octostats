package graphql

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
)

func UserDetails(client *githubv4.Client, user string) interface{} {
	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	err := client.Query(context.Background(), &UserQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	mapUser := UserQuery.User

	jsonResp := map[string]interface{}{
		"username":     mapUser.Login,
		"name":         mapUser.Name,
		"avatar_url":   mapUser.AvatarURL,
		"location":     mapUser.Location,
		"company":      mapUser.Company,
		"email":        mapUser.Email,
		"twitter":      mapUser.TwitterUsername,
		"website_url":  mapUser.WebsiteURL,
		"bio":          mapUser.Bio,
		"status":       mapUser.Status.Emoji + " " + mapUser.Status.Message,
		"created_at":   mapUser.CreatedAt.Time,
		"followers":    mapUser.Followers.TotalCount,
		"following":    mapUser.Following.TotalCount,
		"stargazzers":  mapUser.StarredRepositories.TotalCount,
		"watching":     mapUser.Watching.TotalCount,
		"projects":     mapUser.Projects.TotalCount,
		"packages":     mapUser.Packages.TotalCount,
		"gists":        mapUser.Gists.TotalCount,
		"repositories": mapUser.Repositories.TotalCount,
		"orgs":         mapUser.Organizations.Nodes,
		"sponsoring":   mapUser.SponsorshipsAsSponsor.Nodes,
		"sponsors":     mapUser.SponsorshipsAsMaintainer.Nodes,
	}

	return jsonResp
}
