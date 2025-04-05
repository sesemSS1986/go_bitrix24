package client

func (c *Client) CrmDuplicateFindbycomm(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.duplicate.findbycomm", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
