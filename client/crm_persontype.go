package client

func (c *Client) CrmPersontypeFields(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.persontype.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmPersontypeList(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.persontype.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
