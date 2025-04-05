package client

func (c *Client) ImopenlinesConfigPathGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.config.path.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesConfigGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesConfigListGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.config.list.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesConfigUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.config.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesConfigAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.config.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesConfigDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.config.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
