package project

import (
	"context"

	"github.com/deploifai/sdk-go/api/generated"
)

//Create a new project for the given account.

func (c *Client) Create(ctx context.Context, whereAccount generated.AccountWhereUniqueInput, data generated.CreateProjectInput) (project generated.ProjectFragment, err error) {

	responseData, err := c.options.API.GetGQLClient().CreateProject(ctx, whereAccount, data)
	if err != nil {
		return project, c.options.API.ProcessGQLError(err)
	}

	return *responseData.GetCreateProject(), nil

}
