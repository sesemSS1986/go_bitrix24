package client

func (c *Client) CrmMultifieldFields(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.multifield.fields", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
