package cloud_profile

import (
	"context"
	"github.com/deploifai/sdk-go/api/generated"
)

// Get returns a cloud profile.
func (c *Client) Get(ctx context.Context, where generated.CloudProfileWhereUniqueInput) (cp generated.CloudProfileFragment, err error) {

	data, err := c.options.API.GetGQLClient().GetCloudProfile(ctx, where)
	if err != nil {
		return cp, c.options.API.ProcessGQLError(err)
	}

	return *data.GetCloudProfile(), nil

}
