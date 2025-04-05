package client

func (c *Client) CrmSitebuttonConfigurationGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.sitebutton.configuration.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
