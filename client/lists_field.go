package client

func (c *Client) ListsFieldAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.field.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.field.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.field.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.field.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldTypeGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"lists.field.type.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
