package client

func (c *Client) ForumMessageUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("forum.message.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ForumMessageDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("forum.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
