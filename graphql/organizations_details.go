package graphql

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
)

func OrganizationDetails(client *githubv4.Client, user string) interface{} {
	variables := map[string]interface{}{
		"user": githubv4.String(user),
	}

	err := client.Query(context.Background(), &OrganizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	mapUser := OrganizationQuery.Organization

	jsonResp := map[string]interface{}{
		"username":     mapUser.Login,
		"name":         mapUser.Name,
		"avatar_url":   mapUser.AvatarURL,
		"location":     mapUser.Location,
		"email":        mapUser.Email,
		"twitter":      mapUser.TwitterUsername,
		"website_url":  mapUser.WebsiteURL,
		"description":  mapUser.Description,
		"created_at":   mapUser.CreatedAt.Time,
		"projects":     mapUser.Projects.TotalCount,
		"packages":     mapUser.Packages.TotalCount,
		"repositories": mapUser.Repositories.TotalCount,
		"members":      mapUser.MembersWithRole.TotalCount,
	}

	return jsonResp
}
