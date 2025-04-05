package client

func (c *Client) LogBlogcommentAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "log.blogcomment.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogcommentUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "log.blogcomment.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogBlogcommentDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "log.blogcomment.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
