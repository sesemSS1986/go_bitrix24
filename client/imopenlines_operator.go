package client

func (c *Client) ImopenlinesOperatorAnswer(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.operator.answer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorSkip(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.operator.skip", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorSpam(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.operator.spam", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorTransfer(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.operator.transfer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorFinish(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"imopenlines.operator.finish", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
