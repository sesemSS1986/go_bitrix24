package client

func (c *Client) ForumMessageUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"forum.message.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ForumMessageDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"forum.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
