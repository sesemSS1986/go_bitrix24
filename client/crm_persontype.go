package client

func (c *Client) CrmPersontypeFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.persontype.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmPersontypeList(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.persontype.list", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
