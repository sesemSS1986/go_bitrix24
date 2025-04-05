package client

func (c *Client) ListsFieldAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.field.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.field.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.field.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.field.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ListsFieldTypeGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("lists.field.type.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
