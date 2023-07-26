package project

import (
	"context"

	"github.com/deploifai/sdk-go/api/generated"
)

// Returns a list of cloud profiles for the given account

func (c *Client) List(ctx context.Context, whereAccount generated.AccountWhereUniqueInput, whereProject *generated.ProjectWhereInput) (projects []generated.ProjectFragment, err error) {

	data, err := c.options.API.GetGQLClient().GetProjects(ctx, whereAccount, whereProject)
	if err != nil {
		return nil, c.options.API.ProcessGQLError(err)
	}
	for _, proj := range data.GetProjects() {
		projects = append(projects, *proj)
	}

	return projects, nil
}
