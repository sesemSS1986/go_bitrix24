package client

func (c *Client) CrmPaysystemFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.paysystem.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmPaysystemList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.paysystem.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
