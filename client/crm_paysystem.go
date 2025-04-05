package client

func (c *Client) CrmPaysystemFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.paysystem.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmPaysystemList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.paysystem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
