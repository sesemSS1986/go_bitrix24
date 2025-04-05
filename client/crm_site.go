package client

func (c *Client) CrmSiteFormFill(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.site.form.fill", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) CrmSiteFormUserGet(root bool, url string, p Parameters) (result map[string]interface{}, err error) {
	resp, err := c.Request(root, url, "crm.site.form.user.get", p)
	if err != nil {
		return nil, err
	}
	return resp, err
}
