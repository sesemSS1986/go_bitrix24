package client

func (c *Client) DiskRightsGettasks(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.rights.gettasks", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
