package client

func (c *Client) ImopenlinesOperatorAnswer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.operator.answer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorSkip(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.operator.skip", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorSpam(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.operator.spam", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorTransfer(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.operator.transfer", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImopenlinesOperatorFinish(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.operator.finish", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
