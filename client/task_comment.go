package client

func (c *Client) TaskCommentGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.comment.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.comment.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
