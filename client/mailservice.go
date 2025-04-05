package client

func (c *Client) MailserviceFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mailservice.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mailservice.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mailservice.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceAdd(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mailservice.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceUpdate(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mailservice.update", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) MailserviceDelete(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("mailservice.delete", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
