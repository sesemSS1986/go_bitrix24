package client

func (c *Client) Sonet_groupFeatureAccess(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("sonet_group.feature.access", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
