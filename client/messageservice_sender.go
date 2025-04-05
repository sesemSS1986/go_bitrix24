package client

func (c *Client) MessageserviceSenderAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"messageservice.sender.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MessageserviceSenderDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"messageservice.sender.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MessageserviceSenderList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"messageservice.sender.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
