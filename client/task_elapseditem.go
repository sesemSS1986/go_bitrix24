package client

func (c *Client) TaskElapseditemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskElapseditemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.elapseditem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
