package client

func (c *Client) ImopenlinesNetworkJoin(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.network.join", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesNetworkMessageAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.network.message.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
