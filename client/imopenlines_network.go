package client

func (c *Client) ImopenlinesNetworkJoin(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.network.join", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesNetworkMessageAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.network.message.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
