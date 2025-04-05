package client

func (c *Client) TaskCommentitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCommentitemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.commentitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
