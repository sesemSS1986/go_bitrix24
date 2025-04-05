package client

func (c *Client) TaskCtasklogitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctasklogitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtasklogitemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctasklogitem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
