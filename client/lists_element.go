package client

func (c *Client) ListsElementAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.element.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.element.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.element.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.element.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsElementGetFileUrl(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.element.get.file.url", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
