package client

func (c *Client) ImbotMessageAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.message.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotMessageDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotMessageUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.message.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImbotMessageLike(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "imbot.message.like", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
