package client

func (c *Client) UserHistoryList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.history.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserHistoryFieldsList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"user.history.fields.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
