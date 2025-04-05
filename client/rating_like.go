package client

func (c *Client) LikeList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"like.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LikeReactions(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"like.reactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
