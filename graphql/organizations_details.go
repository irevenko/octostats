package graphql

import (
	"context"
	"log"

	"github.com/shurcooL/githubv4"
)

func OrganizationDetails(client *githubv4.Client, organization string) Organization {
	variables := map[string]interface{}{
		"user": githubv4.String(organization),
	}

	err := client.Query(context.Background(), &OrganizationQuery, variables)
	if err != nil {
		log.Fatal(err)
	}

	return OrganizationQuery.Organization
}
