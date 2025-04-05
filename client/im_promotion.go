package client

func (c *Client) ImPromotionGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.promotion.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImPromotionRead(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"im.promotion.read", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
