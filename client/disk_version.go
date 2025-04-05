package client

func (c *Client) DiskVersionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.version.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
