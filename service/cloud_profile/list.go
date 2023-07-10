package cloud_profile

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

// List returns a list of cloud profiles for the given account.
func (c *Client) List(ctx context.Context, whereAccount generated.AccountWhereUniqueInput, whereCloudProfile *generated.CloudProfileWhereInput) (cloudProfiles []generated.CloudProfileFragment, err error) {

	data, err := c.options.API.GetGQLClient().GetCloudProfiles(ctx, whereAccount, whereCloudProfile)
	if err != nil {
		return nil, c.options.API.ProcessGQLError(err)
	}

	for _, cp := range data.GetCloudProfiles() {
		cloudProfiles = append(cloudProfiles, *cp)
	}

	return cloudProfiles, nil

}
