package client

func (c *Client) UserHistoryList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.history.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) UserHistoryFieldsList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("user.history.fields.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
