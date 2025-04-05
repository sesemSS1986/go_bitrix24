package client

func (c *Client) TaskCtaskcommentitemGetmanifest(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.getmanifest", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemGetlist(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.getlist", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) TaskCtaskcommentitemIsactionallowed(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"task.ctaskcommentitem.isactionallowed", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
