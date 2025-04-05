package client

func (c *Client) MailserviceFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mailservice.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mailservice.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mailservice.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mailservice.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceUpdate(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mailservice.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceDelete(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"mailservice.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
