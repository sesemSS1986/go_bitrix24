package client

func (c *Client) LogBlogcommentAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogcomment.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogcommentUserGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogcomment.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogcommentDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("log.blogcomment.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
