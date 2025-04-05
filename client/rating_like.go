package client

func (c *Client) LikeList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("like.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LikeReactions(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("like.reactions", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
