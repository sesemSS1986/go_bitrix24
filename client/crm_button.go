package client

func (c *Client) CrmButtonList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.button.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmButtonWidgetsGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.button.widgets.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmButtonGuestRegister(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.button.guest.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
