package client

func (c *Client) ImPromotionGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.promotion.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImPromotionRead(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("im.promotion.read", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
