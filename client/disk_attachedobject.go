package client

func (c *Client) DiskAttachedobjectGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"disk.attachedobject.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
