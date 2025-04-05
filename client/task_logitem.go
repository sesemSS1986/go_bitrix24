package client

func (c *Client) TaskLogitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.logitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskLogitemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.logitem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
