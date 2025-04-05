package client

func (c *Client) LogCommentUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.comment.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) LogCommentDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"log.comment.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
