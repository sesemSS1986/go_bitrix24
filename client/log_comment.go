package client

func (c *Client) LogCommentUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.comment.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogCommentDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.comment.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
