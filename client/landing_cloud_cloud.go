package client

func (c *Client) Landing_cloudCloudGetrepository(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing_cloud.cloud.getrepository", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetdemositelist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing_cloud.cloud.getdemositelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetdemopagelist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing_cloud.cloud.getdemopagelist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGeturlpreview(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing_cloud.cloud.geturlpreview", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetappitems(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing_cloud.cloud.getappitems", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Landing_cloudCloudGetappitemmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"landing_cloud.cloud.getappitemmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
