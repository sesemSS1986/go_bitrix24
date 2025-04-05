package client

func (c *Client) ImMessageAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageLike(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.like", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageCommand(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.command", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageShare(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.share", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) ImMessageUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "im.message.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
