package client

func (c *Client) ImopenlinesSessionIntercept(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imopenlines.session.intercept", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
