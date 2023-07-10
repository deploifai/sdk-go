package workspace

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

// List returns a list of workspaces.
// The first element of the list is the user's personal workspace.
// The remaining elements are the user's team workspaces.
func (c *Client) List(ctx context.Context) (workspaces []generated.AccountFragment, err error) {

	data, err := c.options.API.GetGQLClient().GetAccounts(ctx)
	if err != nil {
		return nil, c.options.API.ProcessGQLError(err)
	}

	workspaces = append(workspaces, *data.GetMe().GetAccount())

	for _, teams := range data.GetMe().GetTeams() {
		workspaces = append(workspaces, *teams.GetAccount())
	}

	return workspaces, nil

}
