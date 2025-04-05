package client

func (c *Client) CrmButtonList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.button.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmButtonWidgetsGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.button.widgets.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmButtonGuestRegister(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.button.guest.register", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
