package cloud_profile

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

// Create creates a new cloud profile for the given account.
func (c *Client) Create(ctx context.Context, whereAccount generated.AccountWhereUniqueInput, data generated.CloudProfileCreateInput) (cloudProfile generated.CloudProfileFragment, err error) {

	responseData, err := c.options.API.GetGQLClient().CreateCloudProfile(ctx, whereAccount, data)
	if err != nil {
		return cloudProfile, c.options.API.ProcessGQLError(err)
	}

	return *responseData.GetCreateCloudProfile(), nil

}
