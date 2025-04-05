package client

func (c *Client) CrmSitebuttonConfigurationGet(p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request("crm.sitebutton.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
