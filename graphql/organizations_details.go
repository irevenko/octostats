package graphql

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func OrganizationDetails(client *githubv4.Client, organization string) (org Organization, err error) {
	variables := map[string]interface{}{
		"user": githubv4.String(organization),
	}

	clientErr := client.Query(context.Background(), &OrganizationQuery, variables)
	if clientErr != nil {
		return org, fmt.Errorf("Couldn't get details for organization %s: %w", organization, clientErr)
	}

	return OrganizationQuery.Organization, nil
}
