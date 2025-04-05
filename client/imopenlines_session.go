package client

func (c *Client) ImopenlinesSessionIntercept(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("imopenlines.session.intercept", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
