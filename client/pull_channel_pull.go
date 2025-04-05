package client

func (c *Client) PullConfigGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "pull.config.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullChannelPublicGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "pull.channel.public.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) PullChannelPublicList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "pull.channel.public.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
