package client

func (c *Client) CrmLivefeedmessageAdd(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url,
		"crm.livefeedmessage.add", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
