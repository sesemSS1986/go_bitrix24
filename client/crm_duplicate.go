package client

func (c *Client) CrmDuplicateFindbycomm(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.duplicate.findbycomm", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
