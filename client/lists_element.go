package client

func (c *Client) ListsElementAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "lists.element.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "lists.element.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "lists.element.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "lists.element.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementGetFileUrl(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "lists.element.get.file.url", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
